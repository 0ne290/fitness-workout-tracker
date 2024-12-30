package entities

type Admin struct {
	Guid      uuid.UUID
	CreatedAt time.Time
	Name      string
	Login     string
	Password  string
}

func New(guid uuid.UUID, createdAt time.Time, name string, login string, password string) *Admin, error {
  // Сюда проверОЧКИ
}
