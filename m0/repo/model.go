package repo

const (
	PUBLISHED byte = 0b1 // otherwise `DRAFT`
)

// entity models

type Post struct {
	Model
	Title       string `gorm:"unique;not null"`
	Abstract    string
	Content     string `gorm:"not null"`
	ContentType string
	TOC         string
	PublishedAt int64
	Status      byte
	Tags        uint64
	View
}

type Tag struct {
	Model
	Name string `gorm:"unique"`
	Code uint64 `gorm:"uniqueIndex"`
	View
}

type User struct {
	Model
	Name     string `gorm:"unique"`
	Password string `gorm:"not null"`
	Email    string `gorm:"unique"`
}

// auxiliary models

type Model struct {
	ID        uint64 `gorm:"primarykey"`
	CreatedAt int64
	UpdatedAt int64
}

type View struct {
	ViewPath string `gorm:"unique"`
}
