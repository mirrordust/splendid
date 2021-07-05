package repo

const (
	PUBLISHED byte = 0b1   // otherwise `DRAFT`
	PUBLIC    byte = 0b10  // otherwise `SECRET`
	REPRINT   byte = 0b100 // otherwise `ORIGINAL`
	PREVIEW   byte = 0b1000
	NORMAL         = PUBLISHED | PUBLIC
)

var statusCode = map[string]byte{
	"published": PUBLISHED,
	"public":    PUBLIC,
	"reprint":   REPRINT,
	"preview":   REPRINT,
	"normal":    NORMAL,
}

func Status2Code(status string) byte {
	return statusCode[status]
}

// entity models

type Post struct {
	Model
	Title       string `gorm:"unique"`
	Abstract    string
	Content     string
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
	Password string
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
