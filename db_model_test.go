package gormDBPlus

import (
	"testing"
)

const TestDrive = POSTGRES_DRIVE

type Student struct {
	ID    uint `gorm:"primaryKey"`
	Name  string
	Books []Book `json:"books" gorm:"foreignKey:StudentId"`
}

type Book struct {
	ID        uint
	StudentId uint
	Name      string
}

var TestStudents = []Student{
	{
		ID:   111,
		Name: "111",
	},
	{
		ID:   222,
		Name: "222",
	},
	{
		ID:   333,
		Name: "333",
	},
	{
		ID:   444,
		Name: "444",
	},
	{
		ID:   445,
		Name: "444",
	},
	{
		ID:   446,
		Name: "444",
	},
}

func TestConnector_Save(t *testing.T) {
	c := NewConnector(TestDrive)
	err := c.RawConnection.AutoMigrate(&Student{})
	err = c.RawConnection.AutoMigrate(&Book{})
	if err != nil {
		t.Error("Can not Migrate " + err.Error())
		return
	}
	err = c.RawConnection.Raw("DELETE FROM students where ?", 1).Error
	err = c.RawConnection.Raw("DELETE FROM books where ?", 1).Error
	if err != nil {
		t.Error("Can not Clear " + err.Error())
		return
	}
	err = c.Save(&TestStudents, nil)
	if err != nil {
		t.Error("Can not Created")
	}
	book := Book{Name: "Deneme", StudentId: 222}
	err = c.Save(&book, nil)
	if err != nil {
		t.Error("Can not Created Book")
	}
}

func TestConnector_FetchOne(t *testing.T) {
	c := NewConnector(TestDrive)
	findStudent := Student{ID: TestStudents[2].ID}
	err := c.FetchOne(&findStudent, nil)
	if err != nil {
		t.Error("Can not Find " + err.Error())
		return
	}
	if findStudent.Name != TestStudents[2].Name {
		t.Error("Can not Find " + err.Error())
		return
	}

	findStudent2 := Student{ID: TestStudents[1].ID}
	err = c.FetchOne(&findStudent2, nil, "Books")
	if err != nil {
		t.Error("Can not Find " + err.Error())
		return
	}
	if len(findStudent2.Books) == 0 {
		t.Error("Can not Read Books " + err.Error())
		return
	}
}

func TestConnector_AppendChild(t *testing.T) {
	c := NewConnector(TestDrive)
	appendName := "Adana"
	appendBooks := []Book{
		{Name: "Adana2"},
		{Name: "Adana"},
		{Name: "Adana4"},
		{Name: "Adana5"},
	}
	findStudent := Student{ID: TestStudents[4].ID}
	err := c.AppendChild(&findStudent, "Books", &Book{Name: appendName}, nil)
	err = c.AppendChild(&findStudent, "Books", &appendBooks, nil)
	if err != nil {
		t.Error("Can not Read Books " + err.Error())
		return
	}
	searchStudent := Student{ID: TestStudents[4].ID}
	c.FetchOne(&searchStudent, nil, "Books")
	if searchStudent.Books[0].Name != appendName {
		t.Error("Can not Wrote Books " + err.Error())
		return
	}
	if len(searchStudent.Books) != 5 {
		t.Error("Can not Wrote Bulked Books " + err.Error())
		return
	}
}

func TestConnector_CountChild(t *testing.T) {
	c := NewConnector(TestDrive)
	findStudent := Student{ID: TestStudents[4].ID}
	count := c.CountChild(&findStudent, "Books", nil)
	if count != 5 {
		t.Error("Can not Counted Books ")
		return
	}
}

func TestConnector_ReplaceChild(t *testing.T) {
	c := NewConnector(TestDrive)
	books := []Book{
		{Name: "Bursa1"},
		{Name: "Bursa2"},
		{Name: "Bursa3"},
		{Name: "Bursa4"},
		{Name: "Bursa5"},
		{Name: "Bursa6"},
		{Name: "Bursa7"},
		{Name: "Bursa8"},
	}
	findStudent := Student{ID: TestStudents[4].ID}
	c.ReplaceChild(&findStudent, "Books", &books, nil)
	findStudent2 := Student{ID: TestStudents[4].ID}
	count := c.CountChild(&findStudent2, "Books", nil)
	if count != 8 {
		t.Error("Can not Replaced Books ")
		return
	}
}

func TestConnector_ClearChild(t *testing.T) {
	c := NewConnector(TestDrive)
	findStudent := Student{ID: TestStudents[4].ID}
	c.ClearChild(&findStudent, "Books", nil)
	findStudent2 := Student{ID: TestStudents[4].ID}
	count := c.CountChild(&findStudent2, "Books", nil)
	if count != 0 {
		t.Error("Can not Clear Books ")
		return
	}
}
func TestConnector_Update(t *testing.T) {
	c := NewConnector(TestDrive)
	UpdatedName := "Test123"
	findStudent := Student{ID: TestStudents[1].ID}
	err := c.Update(&findStudent, &Student{Name: UpdatedName}, nil)
	if err != nil {
		t.Error("Can not Deleted")
		return
	}
	findStudentResponse := Student{ID: TestStudents[1].ID}
	c.FetchOne(&findStudentResponse, nil)
	if err != nil {
		t.Error("Can not Find")
		return
	}
	if findStudent.Name != UpdatedName {
		t.Error("Can not Updated")
		return
	}
}

func TestConnector_FetchAll(t *testing.T) {
	c := NewConnector(TestDrive)
	findStudent := Student{Name: "444"}
	students := []Student{}
	c.FetchAll(&students, &findStudent, nil)
	if len(students) != 3 {
		t.Error("Can not Fetch All for Filter")
		return
	}
	allStudents := []Student{}
	c.FetchAll(&allStudents, &Student{}, nil)
	if len(allStudents) != 6 {
		t.Error("Can not Fetch All for All")
		return
	}
}

func TestConnector_FetchOneWithID(t *testing.T) {
	c := NewConnector(TestDrive)
	findStudent := Student{}
	c.FetchOneWithID(&findStudent, int(TestStudents[2].ID), nil)
	if findStudent.Name != TestStudents[2].Name {
		t.Error("Can not Find By Id")
		return
	}
}

func TestConnector_Remove(t *testing.T) {
	c := NewConnector(TestDrive)
	findStudent := Student{ID: TestStudents[2].ID}
	err := c.Remove(&findStudent, nil)
	if err != nil {
		t.Error("Can not Deleted")
		return
	}
	allStudents := []Student{}
	c.FetchAll(&allStudents, &Student{}, nil)
	for _, v := range allStudents {
		err := c.ClearChild(&v, "Books", nil)
		if err != nil {
			t.Error("Can not Deleted Childs : ", err.Error())
			return
		}
	}
	err = c.Remove(&allStudents, nil)
	if err != nil {
		t.Error("Can not Deleted All")
		return
	}
	err = c.RawConnection.Migrator().DropTable(&Student{})
	if err != nil {
		t.Error("Can not Dropped Table")
		return
	}
	err = c.RawConnection.Migrator().DropTable(&Book{})
	if err != nil {
		t.Error("Can not Dropped Table")
		return
	}
}
