package data

import(
	gorm.io/gorm
	gorm.io/driver/sqlite
)

// Migrate program.
func Migrate() {
	db, err := gorm.Open(sqlite.Open("data.sqlite3"), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	db.AutoMigrate(&User{}, &Diary{}, &Target{})
}

