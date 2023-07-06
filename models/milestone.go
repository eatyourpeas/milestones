package models

import (
	"database/sql"
	"strconv"

	_ "github.com/mattn/go-sqlite3"
)

var DB *sql.DB

func ConnectDatabase() error {
	db, err := sql.Open("sqlite3", "./milestones.db")
	if err != nil {
		return err
	}

	DB = db
	return nil
}

type Milestone struct {
	IdMilestone int    `json:"IdMilestone"`
	Milestone   string `json:"Milestone"`
	Description string `json:"Description"`
	Domain      string `json:"Domain"`
	Months      int    `json:"Months"`
}

// func main() {
// 	os.Remove("milestones.db") // I delete the file to avoid duplicated records.
// 	// SQLite is a file based database.

// 	log.Println("Creating milestones.db...")
// 	file, err := os.Create("milestones.db") // Create SQLite file
// 	if err != nil {
// 		log.Fatal(err.Error())
// 	}
// 	file.Close()
// 	log.Println("sqlite-database.db created")

// 	sqliteDatabase, _ := sql.Open("sqlite3", "./milestones.db") // Open the created SQLite File
// 	defer sqliteDatabase.Close()                                // Defer Closing the database
// 	createTable(sqliteDatabase)                                 // Create Database Tables

// 	// INSERT RECORDS
// 	insertMilestone(sqliteDatabase, "0001", "Chin up", "Lifts chin when prone", "Gross Motor", 1)
// 	insertMilestone(sqliteDatabase, "0002", "Turns head", "Turns head in supine postion", "Gross Motor", 1)
// 	insertMilestone(sqliteDatabase, "0003", "Fists", "Hands fisted near face", "Fine Motor", 1)
// 	insertMilestone(sqliteDatabase, "0004", "Voice recognition", "Discriminates parent voice", "Social Skills", 1)
// 	insertMilestone(sqliteDatabase, "0005", "Startles", "Startles to noise", "Hearing", 1)
// 	insertMilestone(sqliteDatabase, "0006", "Stills", "Stills to voice", "Hearing", 1)
// 	insertMilestone(sqliteDatabase, "0007", "Coos", "Coos", "Speech/Language", 2)
// 	insertMilestone(sqliteDatabase, "0008", "Social smile", "Reciprocal smile", "Vision", 2)
// 	insertMilestone(sqliteDatabase, "0009", "Head control", "Holds head when sitting", "Gross Motor", 2)

// 	// DISPLAY INSERTED RECORDS
// 	displayMilestones(sqliteDatabase)
// }

// func createTable(db *sql.DB) {
// 	createMilestoneTableSQL := `CREATE TABLE Milestone (
// 		"IdMilestone" integer NOT NULL PRIMARY KEY AUTOINCREMENT,
// 		"Milestone" TEXT,
// 		"Description" TEXT,
// 		"Domain" TEXT,
// 		"Months" INTEGER
// 	  );` // SQL Statement for Create Table

// 	log.Println("Create Milestone table...")
// 	statement, err := db.Prepare(createMilestoneTableSQL) // Prepare SQL Statement
// 	if err != nil {
// 		log.Fatal(err.Error())
// 	}
// 	statement.Exec() // Execute SQL Statements
// 	log.Println("Milestone table created")
// }

// We are passing db reference connection from main to our method with other parameters
// func insertMilestone(db *sql.DB, IdMilestone string, Milestone string, Description string, Domain string, Months int) {
// 	log.Println("Inserting Milestone record ...")
// 	insertMilestoneSQL := `INSERT INTO Milestone(Milestone,Description,Domain,Months) VALUES (?, ?, ?, ?)`
// 	statement, err := db.Prepare(insertMilestoneSQL) // Prepare statement.
// 	// This is good to avoid SQL injections
// 	if err != nil {
// 		log.Fatalln(err.Error())
// 	}
// 	_, err = statement.Exec(Milestone, Description, Domain, Months)
// 	if err != nil {
// 		log.Fatalln(err.Error())
// 	}
// }

// func displayMilestones(db *sql.DB) {
// 	row, err := db.Query("SELECT * FROM Milestone ORDER BY Months, Milestone")
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	defer row.Close()
// 	for row.Next() { // Iterate and fetch the records from result cursor
// 		var IdMilestone string
// 		var Milestone string
// 		var Description string
// 		var Domain string
// 		var Months int
// 		row.Scan(&IdMilestone, &Milestone, &Description, &Domain, &Months)
// 		log.Println("Milestone: ", Milestone, " ", Description, " ", Domain, " ", Months)
// 	}
// }

func GetMilestones(count int) ([]Milestone, error) {

	rows, err := DB.Query("SELECT IdMilestone, Milestone, Description, Domain, Months from Milestone LIMIT " + strconv.Itoa(count))

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	milestones := make([]Milestone, 0)

	for rows.Next() {
		singleMilestone := Milestone{}
		err = rows.Scan(&singleMilestone.IdMilestone, &singleMilestone.Milestone, &singleMilestone.Description, &singleMilestone.Domain, &singleMilestone.Months)

		if err != nil {
			return nil, err
		}

		milestones = append(milestones, singleMilestone)
	}

	err = rows.Err()

	if err != nil {
		return nil, err
	}

	return milestones, err
}

// func checkErr(err error) {
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// }

// func getMilestones(c *gin.Context) {

// 	milestones, err := models.GetMilestones(10)
// 	checkErr(err)

// 	if milestones == nil {
// 		c.JSON(http.StatusBadRequest, gin.H{"error": "No Records Found"})
// 		return
// 	} else {
// 		c.JSON(http.StatusOK, gin.H{"data": milestones})
// 	}
// }

func GetMilestoneById(IdMilestone string) (Milestone, error) {

	stmt, err := DB.Prepare("SELECT IdMilestone, Milestone, Description, Domain, Months from Milestone WHERE IdMilestone = ?")

	if err != nil {
		return Milestone{}, err
	}

	singleMilestone := Milestone{}

	sqlErr := stmt.QueryRow(IdMilestone).Scan(&singleMilestone.IdMilestone, &singleMilestone.Milestone, &singleMilestone.Description, &singleMilestone.Domain, &singleMilestone.Months)

	if sqlErr != nil {
		if sqlErr == sql.ErrNoRows {
			return Milestone{}, nil
		}
		return Milestone{}, sqlErr
	}
	return singleMilestone, nil
}

func AddMilestone(newMilestone Milestone) (bool, error) {

	tx, err := DB.Begin()
	if err != nil {
		return false, err
	}

	stmt, err := tx.Prepare("INSERT INTO Milestone (IdMilestone, Milestone, Description, Domain, Months) VALUES (?, ?, ?, ?, ?)")

	if err != nil {
		return false, err
	}

	defer stmt.Close()

	_, err = stmt.Exec(newMilestone.IdMilestone, newMilestone.Milestone, newMilestone.Description, newMilestone.Domain, newMilestone.Months)

	if err != nil {
		return false, err
	}

	tx.Commit()

	return true, nil
}

func UpdateMilestone(ourMilestone Milestone, IdMilestone int) (bool, error) {

	tx, err := DB.Begin()
	if err != nil {
		return false, err
	}

	stmt, err := tx.Prepare("UPDATE Milestone SET Milestone = ?, Description = ?, Domain = ?, Months = ? WHERE IdMilestone = ?")

	if err != nil {
		return false, err
	}

	defer stmt.Close()

	_, err = stmt.Exec(ourMilestone.Milestone, ourMilestone.Description, ourMilestone.Domain, ourMilestone.Months, IdMilestone)

	if err != nil {
		return false, err
	}

	tx.Commit()

	return true, nil
}

func DeleteMilestone(IdMilestone string) (bool, error) {

	tx, err := DB.Begin()

	if err != nil {
		return false, err
	}

	stmt, err := DB.Prepare("DELETE from Milestone where IdMilestone = ?")

	if err != nil {
		return false, err
	}

	defer stmt.Close()

	_, err = stmt.Exec(IdMilestone)

	if err != nil {
		return false, err
	}

	tx.Commit()

	return true, nil
}
