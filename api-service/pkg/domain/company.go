package domain

// Company type describes company entity with id, name and company legal form (OOO, ZAO, IP)
type Company struct {
	ID int64
	Name string
	LegalForm string
}