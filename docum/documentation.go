package docum
//this is a golang module

import "fmt"
import "errors"
import "math/rand"

func Hello(name string) (string, error){
	if (name==""){
		return "", errors.New("name empty")
	}
	// If a name was received, return a value that embeds the name
	// in a greeting message.
	return fmt.Sprintf(randomFormat(), name), nil ;

}

// randomFormat returns one of a set of greeting messages. The returned
// message is selected at random.
func randomFormat() string {
    // A slice of message formats.
    formats := []string{
        "Hi, %v. Welcome!",
        "Great to see you, %v!",
        "Hail, %v! Well met!",
    }

    // Return a randomly selected message format by specifying
    // a random index for the slice of formats.
    return formats[rand.Intn(len(formats))]
}

func Hellos(names []string) (map[string]string, error){
	messages := make(map[string]string)

	for _,v := range names{
		message, err := Hello(v)
		if err!=nil {
			return nil, err
		}
		messages[v] = message
	}
	return  messages, nil;

}
