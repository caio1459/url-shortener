package domains

import "time"

type URL struct {
	ID        string     `bson:"_id"`        // slug (ex: AbC123)
	Original  string     `bson:"original"`   // URL longa
	CreatedAt time.Time  `bson:"created_at"` // timestamp
	ExpireAt  *time.Time `bson:"expire_at"`  // expiração opcional
}
