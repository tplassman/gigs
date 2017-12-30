package models

// Fields changed to uppercase
type Gig struct {
	Id          uint
	Name        string
	Description string
	Employer_Id uint
}

func GetGigs() ([]*Gig, error) {
	// Get rows from B
	rows, err := db.Query("SELECT * FROM gigs")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	// Asign DB rows to array of type Gig with scan
	gigs := make([]*Gig, 0)
	for rows.Next() {
		gig := &Gig{}
		err := rows.Scan(
			&gig.Id,
			&gig.Name,
			&gig.Description,
			&gig.Employer_Id,
		) // order matters
		if err != nil {
			return nil, err
		}
		gigs = append(gigs, gig)
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}
	return gigs, nil
}

func GetGig(id uint) (*Gig, error) {
	// Get rows from B
	rows, err := db.Query("SELECT * FROM gigs WHERE id=$1", id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	// Asign DB rows to array of type Gig with scan
	gig := &Gig{}
	for rows.Next() {
		err := rows.Scan(
			&gig.Id,
			&gig.Name,
			&gig.Description,
			&gig.Employer_Id,
		) // order matters
		if err != nil {
			return nil, err
		}
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}
	return gig, nil
}
