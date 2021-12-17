package xorm

import (
	"fmt"
	"testing"

	"xorm.io/xorm"

	"github.com/DATA-DOG/go-sqlmock"
	_ "github.com/go-sql-driver/mysql"
	log "github.com/illidaris/logger"
	. "github.com/smartystreets/goconvey/convey"
)

func TestNewXLogger(t *testing.T) {
	Convey("Setup", t, func() {
		session, mock := getSession()
		repo := NewPersonRepo(session)
		id, name := 1, "John"
		Convey("create a person", func() {
			mock.ExpectExec("INSERT INTO `person`").
				WithArgs(id, name).
				WillReturnResult(sqlmock.NewResult(1, 1))

			err := repo.Create(id, name)
			So(err, ShouldBeNil)
		})

		Convey("create none person", func() {
			mock.ExpectExec("INSERT INTO `person`").
				WithArgs(id, name).
				WillReturnResult(sqlmock.NewResult(0, 0))

			err := repo.Create(id, name)
			So(err, ShouldBeError)
		})

		Reset(func() {
			So(mock.ExpectationsWereMet(), ShouldBeNil)
		})
	})
}

func ExampleNewXLogger() {
	eng, err := xorm.NewEngine("mysql", "root:123@/test?charset=utf8")
	if err != nil {
		panic(err)
	}
	eng.ShowSQL(true)
	// init log core, if not initialized
	log.OnlyConsole()
	// assembly xorm log
	eng.SetLogger(NewXLogger())
}

func getSession() (*xorm.Session, sqlmock.Sqlmock) {
	db, mock, err := sqlmock.New()
	So(err, ShouldBeNil)

	eng, err := xorm.NewEngine("mysql", "root:123@/test?charset=utf8")
	So(err, ShouldBeNil)

	eng.DB().DB = db
	eng.ShowSQL(true)

	log.OnlyConsole()

	eng.SetLogger(NewXLogger())

	return eng.NewSession(), mock
}

type Person struct {
	ID   int    `xorm:"pk id"`
	Name string `xorm:"name"`
}

func (p *Person) TableName() string {
	return "person"
}

type Repository interface {
	Get(id int) (*Person, error)
	Create(id int, name string) error
}

type repo struct {
	session *xorm.Session
}

func NewPersonRepo(session *xorm.Session) Repository {
	return repo{session}
}

func (r repo) Get(id int) (person *Person, err error) {
	person = &Person{ID: id}
	has, err := r.session.Get(person)
	if err != nil {
		return
	}
	if !has {
		err = fmt.Errorf("person[id=%d] not found", id)
		return
	}

	return
}

func (r repo) Create(id int, name string) (err error) {
	person := &Person{ID: id, Name: name}
	affected, err := r.session.Insert(person)
	if err != nil {
		return
	}

	if affected == 0 {
		err = fmt.Errorf("insert err, because of 0 affected")
		return
	}
	return
}
