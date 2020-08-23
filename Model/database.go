package Model

import (
	"Go-Server/Model/utils"
	"encoding/json"
	"io/ioutil"
	"os"
)

type JsonDatabase struct {
	dbFolderName string
	dbName       string
	*Db
}

type Db struct {
	Users      []utils.User
	Authtokens []utils.Authtoken
}

func InitDatabase() *JsonDatabase {
	jdb := JsonDatabase{dbFolderName: "Db folder", dbName: "Db.json"}
	if _, err := os.Stat(jdb.dbFolderName); os.IsNotExist(err) {
		os.Mkdir(jdb.dbFolderName, os.ModePerm)
	}
	if _, err := os.Stat(jdb.dbFolderName + "/" + jdb.dbName); os.IsNotExist(err) {
		dbFile, err := os.Create(jdb.dbFolderName + "/" + jdb.dbName)
		if err != nil {
			panic(err)
		}
		jdb.Db = initDbFromFile(dbFile)
		dbFile.Close()
	} else {
		dbFile, _ := os.Open(jdb.dbFolderName + "/" + jdb.dbName)
		jdb.Db = getDbFromFile(dbFile)
		dbFile.Close()
	}
	return &jdb
}

func initDbFromFile(file *os.File) *Db {
	db := Db{[]utils.User{}, []utils.Authtoken{}}
	if bytArr, err := json.Marshal(db); err == nil {
		if _, err := file.Write(bytArr); err != nil {
			panic(err)
		}
	} else {
		panic(err)
	}
	return &db
}

func getDbFromFile(file *os.File) *Db {
	fileBytes, err := ioutil.ReadFile(file.Name())
	if err != nil {
		panic(err)
	}
	db := &Db{[]utils.User{}, []utils.Authtoken{}}
	if err := json.Unmarshal(fileBytes, db); err != nil {
		panic(err)
	}
	return db
}

func (jdb JsonDatabase) ContainsUser(username string) bool {
	for _, user := range jdb.Users {
		if user.Username == username {
			return true
		}
	}
	return false
}

func (jdb JsonDatabase) ContainAuth(token string) bool {
	for _, authtoken := range jdb.Authtokens {
		if authtoken.Token == token {
			return true
		}
	}
	return false
}

func (jdb JsonDatabase) PostUser(user utils.User) {
	jdb.Db.Users = append(jdb.Db.Users, user)
	jdb.updateDbFile()
}

func (jdb JsonDatabase) PostAuthToken(token string) {
	defer jdb.updateDbFile()
	jdb.Db.Authtokens = append(jdb.Db.Authtokens, utils.Authtoken{Token: token})
}

func (jdb JsonDatabase) updateDbFile() {
	if err := os.Remove(jdb.dbFolderName + "/" + jdb.dbName); err != nil {
		panic(err)
	}
	if dbFile, err := os.Create(jdb.dbFolderName + "/" + jdb.dbName); err == nil {
		if jsonBytes, err := json.Marshal(jdb.Db); err == nil {
			if _, err := dbFile.Write(jsonBytes); err != nil {
				panic(err)
			}
			dbFile.Close()
		} else {
			panic(err)
		}
	} else {
		panic(err)
	}

}

func (jdb JsonDatabase) GetUserByUsername(username string) utils.User {
	for _, user := range jdb.Users {
		if user.Username == username {
			return user
		}
	}
	return utils.User{}
}
