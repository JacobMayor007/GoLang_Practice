// package main

// import "fmt"

// type authenticationInfo struct {
// 	username string
// 	password string
// }

// func (auth authenticationInfo) getBasicAuth() (string, string) {
// 	return auth.username, auth.password
// }

// func main() {
// 	var auth = authenticationInfo{
// 		username: "Jacob",
// 		password: "123123",
// 	}

// 	username, password := (auth.getBasicAuth())

// 	fmt.Printf("Username: %v\nPassword: %v\n", username, password)

// }

// func sendMessage(msg message) (string, int) {
// 	return msg.getMessage(), len(msg.getMessage())
// }

// type message interface {
// 	getMessage() string
// }

// // don't edit below this line

// type birthdayMessage struct {
// 	birthdayTime  int
// 	recipientName string
// }

// func (bm birthdayMessage) getMessage() string {
// 	return fmt.Sprintf("Hi %s, it is your birthday on %v", bm.recipientName, bm.birthdayTime)
// }

// type sendingReport struct {
// 	reportName    string
// 	numberOfSends int
// }

// func (sr sendingReport) getMessage() string {
// 	return fmt.Sprintf(`Your "%s" report is ready. You've sent %v messages.`, sr.reportName, sr.numberOfSends)
// }

// func main() {
// 	var birthMessage = birthdayMessage{
// 		birthdayTime:  24,
// 		recipientName: "Jacob Mary Tapere",
// 	}

// 	var sR = sendingReport{
// 		reportName:    "Jacob Mary Tapere",
// 		numberOfSends: 24,
// 	}

// 	content, cost := sendMessage(birthMessage)
// 	fmt.Println(content, cost)
// 	content2, cost2 := sendMessage(sR)
// 	fmt.Println(content2, cost2)
// }

// type User interface {
// 	Login(email, password string) bool
// 	Welcome() string
// }

// type farmer struct {
// 	farmName string
// 	buyer
// }

// type buyer struct {
// 	email    string
// 	password string
// 	username string
// 	age      uint8
// }

// // farmer implements User
// func (f *farmer) Login(email, password string) bool {
// 	return f.email == email && f.password == password
// }

// func (f *farmer) Welcome() string {
// 	return "Welcome Farmer, " + f.username + " from " + f.farmName
// }

// func (f *farmer) Update(email string) bool {

// 	if f.email == email {
// 		return false
// 	}
// 	if email == "" {
// 		return false
// 	}

// 	f.email = email

// 	return true

// }

// func (b *buyer) Login(email, password string) bool {
// 	return b.email == email && b.password == password
// }

// func (b *buyer) Welcome() string {
// 	return "Welcome Buyer, " + b.username
// }

// func dashboard(u User) {
// 	fmt.Println(u.Welcome())
// }

// func main() {
// 	f := &farmer{
// 		farmName: "Farm 1",
// 		buyer: buyer{
// 			email:    "farmer1@gmail.com",
// 			password: "testfarmer",
// 			username: "Farmer Joe",
// 			age:      24,
// 		},
// 	}

// 	email, password := "", ""

// 	fmt.Print("Enter your email: ")
// 	fmt.Scanln(&email)

// 	fmt.Print("Enter your password: ")
// 	fmt.Scanln(&password)

// 	if f.Login(email, password) {
// 		dashboard(f)
// 	} else {
// 		fmt.Println("Invalid credentials.")
// 	}

// 	newEmail := ""

// 	fmt.Println("Please input new email:")
// 	fmt.Scanln(&newEmail)

// 	fmt.Println(f.Update(newEmail))

// 	fmt.Println("New Email: ", f.email)

// 	sze := reflect.TypeOf(farmer{})
// 	fmt.Printf("The size of the farmer struct: %v\n", sze.Size())

// }

// import "fmt"

// func main() {

// 	result := twoSum([]int{3, 3}, 6)

// 	fmt.Println(result)

// }

// func twoSum(nums []int, target int) []int {
// 	num1, num2 := 0, 0

// 	for i := 0; i < len(nums); i++ {
// 		temp := nums[i]
// 		for j := i + 1; j < len(nums); j++ {
// 			if nums[j]+temp == target {
// 				num1 = i
// 				num2 = j
// 			}
// 		}
// 	}

// 	return []int{num1, num2}
// }

// package main

// import "fmt"

// func (e email) cost() int {
// 	if e.isSubscribed {
// 		return 2
// 	} else {
// 		return 5
// 	}
// }

// func (e email) format() string {
// 	if e.isSubscribed {

// 		return "Subscribed: " + e.body
// 	} else {
// 		return "Not Subscribed: " + e.body
// 	}
// }

// type expense interface {
// 	cost() int
// }

// type formatter interface {
// 	format() string
// }

// type email struct {
// 	isSubscribed bool
// 	body         string
// }

// func getExpense(e expense) {
// 	fmt.Println(e.cost())
// }

// func getFormatter(f formatter) {
// 	fmt.Println(f.format())
// }

// func main() {

// 	em := email{
// 		isSubscribed: false,
// 		body:         "Hello World!",
// 	}

// 	getExpense(em)
// 	getFormatter(em)

// }

// WRONG

// func getExpenseReport(e expense) (string, float64) {

// 	em := email{}
// 	smss := sms{}

// 	switch e {
// 	case e.(email):
// 		return em.toAddress, e.cost()
// 	case e.(sms):
// 		return smss.toPhoneNumber, e.cost()
// 	default:
// 		return "", 0.0
// 	}
// 	// if em, ok := e.(email); ok {
// 	// 	if em.isSubscribed {
// 	// 		return em.toAddress, e.cost()
// 	// 	}

// 	// }
// 	// if sm, ok := e.(sms); ok {
// 	// 	if sm.isSubscribed {
// 	// 		return sm.toPhoneNumber, e.cost()
// 	// 	}
// 	// }
// 	// return "", 0.0
// }

// WRONG

// func getExpenseReport(e expense) (string, float64) {
// 	switch v := e.(type) {
// 	case email:
// 		if v.isSubscribed {
// 			return v.toAddress, e.cost()
// 		}
// 	case sms:
// 		if v.isSubscribed {
// 			return v.toPhoneNumber, e.cost()
// 		}
// 	default:
// 		return "", 0.0
// 	}
// 	return "", 0.0
// }

// don't touch below this line

// type expense interface {
// 	cost() float64
// }

// type email struct {
// 	isSubscribed bool
// 	body         string
// 	toAddress    string
// }

// type sms struct {
// 	isSubscribed  bool
// 	body          string
// 	toPhoneNumber string
// }

// func (e email) cost() float64 {
// 	if !e.isSubscribed {
// 		return float64(len(e.body)) * .05
// 	}
// 	return float64(len(e.body)) * .01
// }

// func (s sms) cost() float64 {
// 	if !s.isSubscribed {
// 		return float64(len(s.body)) * .1
// 	}
// 	return float64(len(s.body)) * .03
// }

// func main() {

// 	em := email{
// 		isSubscribed: true,
// 		body:         "Hello World",
// 		toAddress:    "taperekobskie@gmail.com",
// 	}

// 	ssmss := sms{
// 		isSubscribed:  true,
// 		body:          "Hello World",
// 		toPhoneNumber: "09932237994",
// 	}

// 	choose := "em"

// 	switch choose {
// 	case "sms":
// 		fmt.Println(getExpenseReport(ssmss))
// 	case "em":
// 		fmt.Println(getExpenseReport(em))
// 	}

// }

// package main

// import "fmt"

// type formatter interface {
// 	format() string
// }

// func sendMessage(f formatter) string {

// 	switch v := f.(type) {
// 	case plainText:
// 		return v.message
// 	case bold:
// 		return "**" + v.message + "**"
// 	case code:
// 		return "`" + v.message + "`"
// 	default:
// 		return "Error"
// 	}
// }

// type plainText struct {
// 	message string
// }

// type bold struct {
// 	message string
// }

// type code struct {
// 	message string
// }

// func (pt plainText) format() string {
// 	return pt.message
// }
// func (b bold) format() string {
// 	return b.message
// }
// func (c code) format() string {
// 	return c.message
// }

// func main() {
// 	pt := plainText{
// 		message: "Hello World",
// 	}
// 	b := bold{
// 		message: "Hello World",
// 	}
// 	c := code{
// 		message: "Hello World",
// 	}

// 	fmt.Println(sendMessage(pt))
// 	fmt.Println(sendMessage(b))
// 	fmt.Println(sendMessage(c))
// }

package main

import "fmt"

type notification interface {
	importance() int
}

type directMessage struct {
	senderUsername string
	messageContent string
	priorityLevel  int
	isUrgent       bool
}

type groupMessage struct {
	groupName      string
	messageContent string
	priorityLevel  int
}

type systemAlert struct {
	alertCode      string
	messageContent string
}

func (dm directMessage) importance() int {
	if dm.isUrgent {
		return 50
	}

	return dm.priorityLevel
}

func (gm groupMessage) importance() int {
	return gm.priorityLevel
}

func (sa systemAlert) importance() int {
	return 100
}

func processNotification(n notification) (string, int) {

	switch nn := n.(type) {
	case directMessage:
		return nn.senderUsername, nn.importance()
	case groupMessage:
		return nn.groupName, nn.importance()
	case systemAlert:
		return nn.alertCode, nn.importance()
	default:
		return "Error", 0
	}
}

func main() {
	dm := directMessage{
		senderUsername: "Jacob",
		messageContent: "Hello World",
		priorityLevel:  25,
		isUrgent:       false,
	}
	gm := groupMessage{
		groupName:      "PSS",
		messageContent: "Hello World",
		priorityLevel:  25,
	}

	sa := systemAlert{
		alertCode:      "jk0=921",
		messageContent: "Hello World",
	}

	fmt.Println(processNotification(dm))
	fmt.Println(processNotification(gm))
	fmt.Println(processNotification(sa))

}
