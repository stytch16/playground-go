package NameData

import "fmt"

type panicEmpty struct{}
type panicInvalid struct{}

/* GetPrefixTitle returns Mr. if gender = Masculine, otherwise Mrs. */
func GetPrefixTitle() (prefix string, err error) {
	defer func() {
		switch p := recover(); p {
		case nil:
			/* success! */
		case panicEmpty{}:
			/* Name.gender is empty! -> Send error */
			err = fmt.Errorf("GetPrefixTitle: gender element of Name is empty")
		case panicInvalid{}:
			/* Name.gender is something unrecognizable! -> Send error */
			err = fmt.Errorf("GetPrefixTitle: gender element of Name is not invalid : %s", Name.gender)
		default:
			/* No idea what happened. Resume panic */
			panic(p)
		}
	}()

	switch Name.gender {
	case "Masculine":
		return "Mr.", nil
	case "Feminine":
		return "Ms.", nil
	case "Feminine & Masculine":
		return "Mr/Ms.", nil
	case "":
		panic(panicEmpty{})
	default:
		panic(panicInvalid{})
	}
}

/* GetGender gets the gender of the name, masculine/feminine/... */
func GetGender() (gender string, err error) {
	defer func() {
		switch p := recover(); p {
		case nil:
			/* success */
		case panicEmpty{}:
			err = fmt.Errorf("GetGender: gender element of Name is empty!")
		case panicInvalid{}:
			err = fmt.Errorf("GetGender: gender element of Name is invalid!")
		default:
			panic(p)
		}
	}()
	switch Name.gender {
	case "Masculine":
		return "masculine", nil
	case "Feminine":
		return "feminine", nil
	case "Feminine & Masculine":
		return "Mr/Ms.", nil
	case "":
		panic(panicEmpty{})
	default:
		panic(panicInvalid{})
	}
}

/* GetMeaning gets the meaning of the name in string form */
func GetMeaning() (meaning string, err error) {
	defer func() {
		switch p := recover(); p {
		case nil:
			/* success */
		case panicEmpty{}:
			err = fmt.Errorf("GetMeaning: meaning element of Name is empty!")
		default:
			panic(p)
		}
	}()
	switch Name.meaning {
	case "":
		panic(panicEmpty{})
	default:
		return Name.meaning, nil
	}
}

/* GetUsage ... */
func GetUsage() (usage string, err error) {
	defer func() {
		switch p := recover(); p {
		case nil:
			/* success */
		case panicEmpty{}:
			err = fmt.Errorf("GetUsage: usage element is empty!")
		default:
			panic(p)
		}
	}()
	switch Name.usage {
	case "":
		panic(panicEmpty{})
	default:
		return Name.usage, err
	}
}
