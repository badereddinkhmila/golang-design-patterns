package creational

type Address struct {
	Housenumber *int
	Street      *string
	City        string
	Postcode    int
	Country     string
}

type User struct {
	Firstname *string
	Lastname  *string
	Age       *int
	Email     string
	Password  string
	Address   *Address
}

/*********************************************************************/
/*************************** All Fields ******************************/
/*********************************************************************/

func NewUser(firstname, lastname *string, password, email string, age *int, address *Address) *User {
	return &User{
		Firstname: firstname,
		Lastname:  lastname,
		Email:     email,
		Password:  password,
		Age:       age,
		Address:   address,
	}
}

/*********************************************************************/
/*************************** Build object ****************************/
/*********************************************************************/

func InitUser(email, pwd string) *User {
	return &User{Email: email, Password: pwd}
}

func (u *User) setFirstname(firstname string) *User {
	u.Firstname = &firstname
	return u
}

func (u *User) setLastname(firstname string) *User {
	u.Firstname = &firstname
	return u
}

func (u *User) setAge(age int) *User {
	u.Age = &age
	return u
}

func (u *User) setAddress(address *Address) *User {
	u.Address = address
	return u
}
