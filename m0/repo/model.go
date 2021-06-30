package repo

type StatusType byte

const (
	DRAFT     StatusType = 0
	PUBLISHED StatusType = 0b1
	PREVIEW   StatusType = 0b10
	PUBLIC    StatusType = 0b100  // otherwise `SECRET`
	REPRINT   StatusType = 0b1000 // otherwise `ORIGINAL`
)

// max number of tags = 64
type TagType uint64

// entity models
type Post struct {
	Model
	Title       string
	Abstract    string
	Content     string
	ContentType string
	TOC         string
	Status      StatusType
	Tags        TagType
	View
}

type Tag struct {
	Model
	Name string
	Code TagType
	View
}

type User struct {
	Model
	Name     string
	Password string
	Email    string
}

// auxiliary models
type Model struct {
	ID        uint64 `gorm:"primarykey"`
	CreatedAt int64
	UpdatedAt int64
}

type View struct {
	ViewPath string
}
