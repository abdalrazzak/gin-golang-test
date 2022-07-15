type User struct {
	ID           uint 	 `json:"id"`
	Email        *string `json:"email"`
	Password     *string `json:"password"`
	Age          uint8   `json:"age"` 
	CreatedAt    time.Time
	UpdatedAt    time.Time
  }