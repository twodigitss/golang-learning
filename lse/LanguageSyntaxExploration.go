package lse
import "fmt"

//## VARIABLES AND CONSTS
func Vars(){
	var a int = 10;
	var b = 20;
	c := 30; // short declaration (only inside functions)

	var x, y = 1, 2;
	var a_, b_ = "12", 12;

	const Pi = 3.14
	fmt.Print(a,b,c,x,y,a_,b_)

}

//## DATA TYPES
func dataTypes(){
	var arr [3]int = [3]int{1, 2, 3} // Array (fixed size)
	var s = []int{1, 2, 3} // Slice (dynamic)
	var m = map[string]int{ // Map (hash table)
		"a": 1,
		"b": 2,
	};
	fmt.Print(arr,s,m)

}


//## CONTROL FLOW
func controlFlow() {
	//already existing variable
	const b = 5;
	if b > 10 {
    	fmt.Println("big")
    }  else {
    	fmt.Println("small")
    }

    //with var declaration
    if v := 5; v > 0 {
        fmt.Println(v)
    }

    //switch
    switch b {
    case 1:
        fmt.Println("one")
    case 2, 3:
        fmt.Println("two or three")
    default:
        fmt.Println("other")
    }
}

//LOOPS
func forLoop(){
	//common loop
	for i := 0; i<10; i++ {
		fmt.Print(i)
	}

	//range
	for i := range 10 {
		fmt.Print(i)
	}

	//for each element in a array
	nums := []int{1, 2, 3}
	for i, v := range nums {
		fmt.Println(i, v)
	}

	//while kinda
	for i:=0; i<10; {
		i++;
	}

	//infinite
	for {
		break;
	}

}

//## FUNCTIONS
func add(a int, b int) int {
    return a + b
}

//short
func addSimple(a, b int) int {
    return a + b
}

//multiple return errors
func divide(a, b int) (int, error) {
    if b == 0 {
        return 0, fmt.Errorf("division by zero")
    }
    return a / b, nil
}

//nested functions
func superior(){
	nestedFunc := func(name string){
		fmt.Print(name)
	}

	nestedFunc("lil peep")
}


//## Structs

//the structure itself
type user struct{
	name string
	age uint8
}

//functions associated to struts
func (u user) Greet() string {
	// WARN: u is a copy of the original struct
	// Any modification inside the method does NOT affect the original
	return "Hello " + u.name
}

func (u *user) birthday(){
	// WARN: u is a reference to the original struct
	// Changes persist outside the method
	fmt.Print("Hppy bday!");
	u.age++;
}


func structs(){

	var me user = user{"juan", 18}
	me2 := user{"juanjo", 18}

	me.Greet()

	fmt.Print(me,me2)


}


//## INTERFACES
type PaymentMethod interface {
    Pay(amount float64) string
}

type Cash struct {
    balance float64
}

func (c Cash) Pay(amount float64) string {
    if c.balance < amount {
        return "not enough cash"
    }
    c.balance -= amount
    return "paid with cash"
}

func interfaceUsage(){
	var c Cash = Cash{balance: 100}

	var p PaymentMethod = c
	result := p.Pay(50);

	fmt.Println(result)
	fmt.Println(c.balance) // ???


}


//## ERROR HANDLING
func err(){
	var result, err = divide(10, 0)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Print(result)

}


//##POINTERS
func Pointers(){
	x := 10
	p := &x
	fmt.Println(*p) // dereference
}

//##CONCURRENCY

//go routine
func async(){
	go func() {
		fmt.Println("running async")
	}()
}

//channels
func channelss(){
	ch := make(chan int)

	go func() {
		ch <- 42
	}()

	value := <-ch
	fmt.Print(value)

}


//##DEFERS
func deferr(){
	defer fmt.Println("cleanup") //then after all just finished, this one
	fmt.Println("main") //this will execute first
}
