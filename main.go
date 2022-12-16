package main

import (
	"database/sql"
	"fmt"
	"io"
	"io/ioutil"
	"path/filepath"

	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"

	// "log"
	"os"
	"strings"

	"github.com/360EntSecGroup-Skylar/excelize"

	"encoding/json"
	"log"
	"net/http"
	"strconv"
	"time"
)

func init() {
	if len(os.Args) > 1 {
		os.Setenv("ENV_FILE", filepath.Join("./config", os.Args[1]))
	} else {
		os.Setenv("ENV_FILE", filepath.Join("./config", ".env"))
	}

	err := godotenv.Load(os.Getenv("ENV_FILE"))
	if err != nil {
		panic(err.Error())
	}

	log.SetFlags(log.LstdFlags | log.Lshortfile)
}

// func main(){

// resp,err:=http.Get("https://gitlab.smartfren.com/api/v4/projects?private_token="+os.Getenv("PRIVATE_TOKEN")+"&per_page=100&page=7")

// 	if err != nil{
// 		log.Fatalln(err)
// 	}

// 	// body,err :=ioutil.ReadAll(resp.Body)

// 	defer resp.Body.Close()

// 	var result map[string]interface{}

// 	json.NewDecoder(resp.Body).Decode(&result)

// 	fmt.Println(result["name"])

// 	// if err!=nil{
// 	// 	log.Fatalln(err)
// 	// }

// 	// log.Println(string(body))
//     fmt.Println("a")
// }

type RepositoryOwner struct {
	Username string `json:"username"`
	state    string `json:"state"`
}

type RepositoryNameSpace struct {
	Name string `json:"name"`
}

type Repository struct {
	ID          int                 `json:"id"`
	Owner       RepositoryOwner     `json:"owner"`
	Name        string              `json:"name"`
	WebUrl      string              `json:"web_url"`
	Description string              `json:"description"`
	NameSpace   RepositoryNameSpace `json:"namespace"`
	CreatorID   int                 `json:"creator_id"`
	CreateAt    string              `json:"created_at"`
}

type Member struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	Username    string `json:"username"`
	AccessLevel int    `json:"access_level"`
	State       string `json:"state"`
}

type Creator struct {
	Name     string `json:"name"`
	Username string `json:"username"`
}

type AllUserPage struct {
	Page     int
	ListUser []*Users
}

type Users struct {
	ID             int    `json:"id"`
	Name           string `json:"name"`
	Username       string `json:"username"`
	State          string `json:"state"`
	LastActivityOn string `json:"last_activity_on"`
	CurrSignIn     string `json:"current_sign_in_at"`
	Email          string `json:"email"`
	IsAdmin        bool   `json:"is_admin"`
	CreatedAt      string `json:"created_at"`
	Datedif        string
}

type DataKaryawan struct {
	Nik           string
	Nama          string
	Email         string
	Join          string
	Groups        string
	SuperiorName  string
	SuperiorEmail string
}

type GitlabList struct {
	NameProject string
	UrlProject  string
	Owner       string
	Member      string
	Superior1   string
	Superior2   string
}

// func main() {

// 	license()

// 	os.Exit(71)

// 	errAll, allKaryawan := getKaryawan()

// 	if errAll != nil {
// 		log.Println("failed get data karyawan")
// 		os.Exit(71)
// 	}

// 	// fmt.Println(len(allKaryawan))
// 	// os.Exit(71)

// 	var arrStr []string

// 	arrStr = append(arrStr, "project_name")
// 	arrStr = append(arrStr, "|")
// 	// arrStr = append(arrStr, "project_description")
// 	// arrStr = append(arrStr, "|")
// 	arrStr = append(arrStr, "url_project")
// 	arrStr = append(arrStr, "|")
// 	arrStr = append(arrStr, "owner")
// 	arrStr = append(arrStr, "|")
// 	arrStr = append(arrStr, "member")
// 	arrStr = append(arrStr, "|")
// 	arrStr = append(arrStr, "superior1")
// 	arrStr = append(arrStr, "|")
// 	arrStr = append(arrStr, "superior2")
// 	arrStr = append(arrStr, "\n")

// 	// for i := 0; i <= 6; i++ {
// 	i:=0

// 	gitlabList := make([]*GitlabList,0)

// 	for{
// 		pageIdx := i + 1
// 		fmt.Println(pageIdx)
// 		page := strconv.Itoa(pageIdx)
// res, err := http.Get("https://gitlab.smartfren.com/api/v4/projects?private_token="+os.Getenv("PRIVATE_TOKEN")+"&per_page=100&page=" + page + "")

// 		// fmt.Println("page ke ", page)

// 		if err != nil {
// 			panic(err.Error())
// 		}

// 		body, err := ioutil.ReadAll(res.Body)

// 		if err != nil {
// 			panic(err.Error())
// 		}

// 		structJSON := make([]Repository, 0)

// 		json.Unmarshal(body, &structJSON)

// 		if len(structJSON) == 0 {
// 			break
// 		}

// 		idx := 0

// 		for idx < len(structJSON) {

// 			bk := new(GitlabList)

// 			idProject := structJSON[idx].ID

// 			nameProject := structJSON[idx].Name

// 			// descProject := structJSON[idx].Description

// 			urlProject := structJSON[idx].WebUrl

// 			creatorID := structJSON[idx].CreatorID

// 			// strMember := ""

// 			// if idProject == 1257{
// 			// 	strMember = getMember(idProject)
// 			// 	os.Exit(71)
// 			// }

// 			// if idProject != 1247 {
// 			// 	idx++
// 			// 	continue
// 			// }
// 			strMember,newStructMember := getMember(idProject)

// 			// os.Exit(71)

// 			var sliceMember []string

// 			for _,rangeMember := range newStructMember {
// 				if rangeMember.Username == strMember{
// 					continue
// 				}
// 				sliceMember = append(sliceMember,rangeMember.Username)
// 			}

// 			allMember := strings.Join(sliceMember,",")

// 			if strMember == "unknown" {
// 				idx++
// 				continue
// 				strMember = getCreator(creatorID)
// 			}

// 			// errDetailOwner,detailOwner := readDb(strMember+"@smartfren.com")

// 			// superiorData := detailOwner.SuperiorEmail

// 			// if errDetailOwner != nil {
// 			// 	superiorData = "unknown"
// 			// }

// 			ownerEmail := strMember+"@smartfren.com"

// 			superior1,superior2 := getSuperior(ownerEmail,allKaryawan)

// 			bk.NameProject = nameProject
// 			bk.UrlProject = urlProject
// 			bk.Owner = strMember
// 			bk.Member = allMember
// 			bk.Superior1 = superior1
// 			bk.Superior2 = superior2

// 			gitlabList = append(gitlabList,bk)

// 			arrStr = append(arrStr, nameProject)
// 			arrStr = append(arrStr, "|")
// 			// arrStr = append(arrStr, descProject)
// 			// arrStr = append(arrStr, "|")
// 			arrStr = append(arrStr, urlProject)
// 			arrStr = append(arrStr, "|")
// 			arrStr = append(arrStr, strMember)
// 			arrStr = append(arrStr, "|")
// 			arrStr = append(arrStr, allMember)
// 			arrStr = append(arrStr, "|")
// 			arrStr = append(arrStr, superior1)
// 			arrStr = append(arrStr, "|")
// 			arrStr = append(arrStr, superior2)
// 			arrStr = append(arrStr, "\n")

// 			idx++
// 		}
// 		i++
// 	}

// 	errExcel := writeToExcel(gitlabList)

// 	if errExcel != nil {
// 		log.Println("error when read data to excel")
// 		os.Exit(71)
// 	}

// 	strWrite := strings.Join(arrStr, "")

// 	filename := "gitlab-project.csv"

// 	tes, errWri := writeToFile(filename, strWrite)

// 	if errWri != nil {
// 		log.Fatalln(errWri)
// 	}

// 	fmt.Println(tes)
// }

func main() {

	license()
	a := getAllRepo()
	fmt.Println(a)
	// os.Exit(71)
	users := getUsers()

	writeUser := writeAllUser(users)

	fmt.Println(writeUser)

	// return

	// os.Exit(71)

	// for _,rangeUsers := range users{
	// 	fmt.Println(rangeUsers.Page,"===========================================")
	// 	for _,rangeList := range rangeUsers.ListUser{
	// 		fmt.Println(rangeList.Name)
	// 	}
	// }

	user90D := getInactiveUser(users)

	if user90D != nil {
		log.Println(user90D)
	}

	// for _,rangeInactive := range user90D{
	// 	fmt.Println(rangeInactive.Name,rangeInactive.Datedif,rangeInactive.ID)
	// }
}

func writeAllUser(users []*AllUserPage) error {
	allUser := make([]*Users, 0)
	for _, rangeAll := range users {
		for _, detailUser := range rangeAll.ListUser {
			pUser := new(Users)
			pUser.Username = detailUser.Username
			pUser.State = detailUser.State
			pUser.Email = detailUser.Email
			pUser.IsAdmin = detailUser.IsAdmin
			pUser.LastActivityOn = detailUser.LastActivityOn
			pUser.CreatedAt = detailUser.CreatedAt
			allUser = append(allUser, pUser)
		}
	}

	var arrWrite []string
	arrWrite = append(arrWrite, "no.")
	arrWrite = append(arrWrite, "|")
	arrWrite = append(arrWrite, "username")
	arrWrite = append(arrWrite, "|")
	arrWrite = append(arrWrite, "state")
	arrWrite = append(arrWrite, "|")
	arrWrite = append(arrWrite, "email")
	arrWrite = append(arrWrite, "|")
	arrWrite = append(arrWrite, "is_admin")
	arrWrite = append(arrWrite, "|")
	arrWrite = append(arrWrite, "last_activity_on")
	arrWrite = append(arrWrite, "|")
	arrWrite = append(arrWrite, "created_at")
	arrWrite = append(arrWrite, "|")
	arrWrite = append(arrWrite, "sub_group")
	arrWrite = append(arrWrite, "\n")

	no := 1
	for _, rangeUser := range allUser {
		isAdmin := "no"
		if rangeUser.IsAdmin {
			isAdmin = "yes"
		}

		_, subGroup := getSubGroup(rangeUser.Email)

		arrCreatedAt := strings.Split(rangeUser.CreatedAt, "T")
		strCreatedAt := arrCreatedAt[0]

		arrWrite = append(arrWrite, strconv.Itoa(no))
		arrWrite = append(arrWrite, "|")
		arrWrite = append(arrWrite, rangeUser.Username)
		arrWrite = append(arrWrite, "|")
		arrWrite = append(arrWrite, rangeUser.State)
		arrWrite = append(arrWrite, "|")
		arrWrite = append(arrWrite, rangeUser.Email)
		arrWrite = append(arrWrite, "|")
		arrWrite = append(arrWrite, isAdmin)
		arrWrite = append(arrWrite, "|")
		arrWrite = append(arrWrite, rangeUser.LastActivityOn)
		arrWrite = append(arrWrite, "|")
		arrWrite = append(arrWrite, strCreatedAt)
		arrWrite = append(arrWrite, "|")
		arrWrite = append(arrWrite, subGroup)
		arrWrite = append(arrWrite, "\n")
		no++
	}

	strWrite := strings.Join(arrWrite, "")

	filename := "gitlab-all-user.csv"

	tes, errWri := writeToFile(filename, strWrite)

	if errWri != nil {
		return errWri
	}

	log.Println(tes + " already generated")

	return errWri
}

func getInactiveUser(users []*AllUserPage) error {
	inactiveUser := make([]*Users, 0)

	y, m, d := getCurrentTime()

	for i, rangeUser := range users {
		fmt.Println(i)
		for _, detailUser := range rangeUser.ListUser {
			pUser := new(Users)

			pUser.ID = detailUser.ID
			pUser.Name = detailUser.Name
			pUser.Username = detailUser.Username
			pUser.State = detailUser.State
			pUser.CurrSignIn = detailUser.CurrSignIn
			pUser.Email = detailUser.Email
			pUser.IsAdmin = detailUser.IsAdmin
			pUser.Datedif = detailUser.Datedif

			if len(detailUser.CurrSignIn) > 0 {
				arrLastAct := strings.Split(detailUser.CurrSignIn, "-")
				yLast := convertToInt(arrLastAct[0])
				mLast := convertToInt(arrLastAct[1])
				dLast := convertToInt(arrLastAct[2])
				Time := time.Date(yLast, time.Month(mLast), dLast, 0, 0, 0, 0, time.UTC)
				t2 := Time.AddDate(0, 0, 0)

				yDif, _ := strconv.Atoi(t2.Format("2006"))
				mDif, _ := strconv.Atoi(t2.Format("01"))
				dDif, _ := strconv.Atoi(t2.Format("02"))

				timeNow := Date(y, m, d)

				timeLast := Date(yDif, mDif, dDif)

				dateDif := int(timeNow.Sub(timeLast).Hours() / 24)

				if dateDif >= 90 {
					pUser.Datedif = strconv.Itoa(dateDif)
					inactiveUser = append(inactiveUser, pUser)
				}
			} else if len(detailUser.CurrSignIn) == 0 {
				pUser.CurrSignIn = "never access"
				pUser.Datedif = "never access"
				inactiveUser = append(inactiveUser, pUser)
			}
		}
	}

	var arrWrite []string
	arrWrite = append(arrWrite, "no.")
	arrWrite = append(arrWrite, "|")
	arrWrite = append(arrWrite, "id_gitlab")
	arrWrite = append(arrWrite, "|")
	arrWrite = append(arrWrite, "username")
	arrWrite = append(arrWrite, "|")
	arrWrite = append(arrWrite, "state")
	arrWrite = append(arrWrite, "|")
	arrWrite = append(arrWrite, "email")
	arrWrite = append(arrWrite, "|")
	arrWrite = append(arrWrite, "last_sign_in")
	arrWrite = append(arrWrite, "|")
	arrWrite = append(arrWrite, "not used for (days)")
	arrWrite = append(arrWrite, "\n")

	fmt.Println(len(inactiveUser))
	no := 1
	for _, rangeUser := range inactiveUser {
		newCurrSignIn := rangeUser.CurrSignIn
		if rangeUser.CurrSignIn != "never access" {
			currSignInFormat := strings.Split(rangeUser.CurrSignIn, "-")
			ySignIn := currSignInFormat[0]
			mSignIn := currSignInFormat[1]
			dSignIn := currSignInFormat[2][0:2]
			newCurrSignIn = ySignIn + "-" + mSignIn + "-" + dSignIn
			fmt.Println(newCurrSignIn)
		}
		arrWrite = append(arrWrite, strconv.Itoa(no))
		arrWrite = append(arrWrite, "|")
		arrWrite = append(arrWrite, strconv.Itoa(rangeUser.ID))
		arrWrite = append(arrWrite, "|")
		arrWrite = append(arrWrite, rangeUser.Username)
		arrWrite = append(arrWrite, "|")
		arrWrite = append(arrWrite, rangeUser.State)
		arrWrite = append(arrWrite, "|")
		arrWrite = append(arrWrite, rangeUser.Email)
		arrWrite = append(arrWrite, "|")
		arrWrite = append(arrWrite, newCurrSignIn)
		arrWrite = append(arrWrite, "|")
		arrWrite = append(arrWrite, rangeUser.Datedif)
		arrWrite = append(arrWrite, "\n")
		no++
	}

	strWrite := strings.Join(arrWrite, "")

	filename := "gitlab-inactive-user_" + strconv.Itoa(y) + strconv.Itoa(m) + strconv.Itoa(d) + ".csv"

	tes, errWri := writeToFile(filename, strWrite)

	if errWri != nil {
		return errWri
	}

	log.Println(tes + " already generated")

	return errWri
}

func getCurrentTime() (int, int, int) {
	currentTime := time.Now()
	y, _ := strconv.Atoi(currentTime.Format("2006"))
	m, _ := strconv.Atoi(currentTime.Format("01"))
	d, _ := strconv.Atoi(currentTime.Format("02"))
	return y, m, d
}

func Date(year, month, day int) time.Time {
	return time.Date(year, time.Month(month), day, 0, 0, 0, 0, time.UTC)
}

func convertToInt(param string) int {
	result, errConv := strconv.Atoi(param)
	if errConv != nil {
		result = 0
	}
	return result
}

func getMember(idProject int) (string, []*Member) {
	stridProject := strconv.Itoa(idProject)

	resp, err1 := http.Get("https://gitlab.smartfren.com/api/v4/projects/" + stridProject + "/members/all?private_token=" + os.Getenv("PRIVATE_TOKEN"))

	if err1 != nil {
		panic(err1.Error())
	}

	body1, err2 := ioutil.ReadAll(resp.Body)

	if err2 != nil {
		panic(err2.Error())
	}

	structMember := make([]Member, 0)

	newStructMember := make([]*Member, 0)

	json.Unmarshal(body1, &structMember)

	// fmt.Println(len(structMember))

	idx := 0

	maxAcc := 0

	maxUsername := "unknown"

	for idx < len(structMember) {

		bk := new(Member)

		username := structMember[idx].Username

		state := structMember[idx].State

		// fmt.Println(username,state)

		if state == "ldap_blocked" || state == "blocked" {
			idx++
			continue
		}

		bk.ID = structMember[idx].ID
		bk.Name = structMember[idx].Name
		bk.Username = structMember[idx].Username
		bk.AccessLevel = structMember[idx].AccessLevel
		bk.State = structMember[idx].State

		newStructMember = append(newStructMember, bk)
		// ldapName := structMember[idx].Username

		accessLevel := structMember[idx].AccessLevel

		if len(structMember) > 1 && accessLevel > maxAcc {
			// if len(structMember) > 1 && accessLevel > maxAcc && username != "nugroho.suwito" {
			maxUsername = username
			maxAcc = accessLevel
		} else if len(structMember) == 1 && accessLevel > maxAcc {
			maxUsername = username
			maxAcc = accessLevel
		}

		// fmt.Println(maxUsername)
		idx++
	}

	return maxUsername, newStructMember

}

func getCreator(id int) string {
	strID := strconv.Itoa(id)

	resp, err1 := http.Get("https://gitlab.smartfren.com/api/v4/users/" + strID + "?private_token=" + os.Getenv("PRIVATE_TOKEN"))

	if err1 != nil {
		panic(err1.Error())
	}

	body1, err2 := ioutil.ReadAll(resp.Body)

	if err2 != nil {
		panic(err2.Error())
	}

	structCreator := new(Creator)

	json.Unmarshal(body1, &structCreator)

	return structCreator.Name

}

func license() {
	errKaryawan, allKaryawan := getKaryawan()

	if errKaryawan != nil {
		log.Println(errKaryawan)
		os.Exit(71)
	}

	listUser := getUsers()

	fmt.Println(len(allKaryawan))

	var arrStr []string
	no := 1

	arrStr = append(arrStr, "no")
	arrStr = append(arrStr, "|")
	arrStr = append(arrStr, "name")
	arrStr = append(arrStr, "|")
	arrStr = append(arrStr, "status")
	arrStr = append(arrStr, "|")
	arrStr = append(arrStr, "email")
	arrStr = append(arrStr, "|")
	arrStr = append(arrStr, "groups")
	arrStr = append(arrStr, "\n")

	for _, rangeList := range listUser {
		for i, rangeUser := range rangeList.ListUser {
			count := 0
			fmt.Println(i)
			for _, rangeKaryawan := range allKaryawan {
				if rangeUser.Email == strings.ToLower(rangeKaryawan.Email) {
					count++
					arrStr = append(arrStr, strconv.Itoa(no))
					arrStr = append(arrStr, "|")
					arrStr = append(arrStr, rangeKaryawan.Nama)
					arrStr = append(arrStr, "|")
					arrStr = append(arrStr, rangeUser.State)
					arrStr = append(arrStr, "|")
					arrStr = append(arrStr, rangeKaryawan.Email)
					arrStr = append(arrStr, "|")
					arrStr = append(arrStr, rangeKaryawan.Groups)
					arrStr = append(arrStr, "\n")
				}
			}
			if count == 0 {
				arrStr = append(arrStr, strconv.Itoa(no))
				arrStr = append(arrStr, "|")
				arrStr = append(arrStr, rangeUser.Name)
				arrStr = append(arrStr, "|")
				arrStr = append(arrStr, rangeUser.State)
				arrStr = append(arrStr, "|")
				arrStr = append(arrStr, rangeUser.Email)
				arrStr = append(arrStr, "|")
				arrStr = append(arrStr, "undefined")
				arrStr = append(arrStr, "\n")
			}
			no++
		}
	}

	strWrite := strings.Join(arrStr, "")

	filename := "license.csv"

	tes, errWri := writeToFile(filename, strWrite)

	if errWri != nil {
		log.Fatalln(errWri)
	}

	fmt.Println(tes)

	fmt.Println(len(listUser))
}

func getUsers() []*AllUserPage {

	allUser := make([]*AllUserPage, 0)

	i := 1

	tes := 0

	for {
		fmt.Println(i)
		detailUser := new(AllUserPage)

		listUser := make([]*Users, 0)

		resp, err := http.Get("https://gitlab.smartfren.com/api/v4/users?private_token=" + os.Getenv("PRIVATE_TOKEN") + "&per_page=100&page=" + strconv.Itoa(i))

		if err != nil {
			panic(err.Error())
		}

		body, errRead := ioutil.ReadAll(resp.Body)

		if errRead != nil {
			panic(errRead.Error())
		}

		json.Unmarshal(body, &listUser)

		tes += len(listUser)

		if len(listUser) == 0 {
			break
		}

		detailUser.Page = i
		detailUser.ListUser = listUser

		allUser = append(allUser, detailUser)

		i++

		// fmt.Println("plus")
	}
	// fmt.Println(len(allUser),"cek",tes)
	return allUser
}

func getAllRepo() int {
	var arrStr []string
	arrStr = append(arrStr, "project_id")
	arrStr = append(arrStr, "|")
	arrStr = append(arrStr, "project_name")
	arrStr = append(arrStr, "|")
	// arrStr = append(arrStr, "project_description")
	// arrStr = append(arrStr, "|")
	arrStr = append(arrStr, "url_project")
	arrStr = append(arrStr, "|")
	arrStr = append(arrStr, "created_date")
	arrStr = append(arrStr, "|")
	arrStr = append(arrStr, "owner")
	arrStr = append(arrStr, "\n")

	i := 1
	a := 0

	for {
		fmt.Println(i)

		res, err := http.Get("https://gitlab.smartfren.com/api/v4/projects?private_token=" + os.Getenv("PRIVATE_TOKEN") + "&per_page=100&page=" + strconv.Itoa(i))

		if err != nil {
			panic(err.Error())
		}

		body, err := ioutil.ReadAll(res.Body)

		if err != nil {
			panic(err.Error())
		}

		structJSON := make([]Repository, 0)

		json.Unmarshal(body, &structJSON)

		idx := 0

		if len(structJSON) == 0 {
			break
		}

		// fmt.Println(i,structJSON[0])
		// continue

		for idx < len(structJSON) {

			idProject := structJSON[idx].ID

			dateCreated := structJSON[idx].CreateAt

			arrDate := strings.Split(dateCreated, "T")

			onlyDate := arrDate[0]

			arrayDate := strings.Split(onlyDate, "-")

			month := convertToInt(arrayDate[1])

			year := convertToInt(arrayDate[0])

			// if year < 2020 {
			// 	continue
			// }

			// if (year < 2020) || (year == 2020 && month < 7){
			// 	fmt.Println(year,month,idProject,"cek")
			// 	idx++
			// 	continue
			// }

			fmt.Println(year, month, idProject)

			// idProject := structJSON[idx].ID

			nameProject := structJSON[idx].Name

			// descProject := structJSON[idx].Description

			urlProject := structJSON[idx].WebUrl

			creatorID := structJSON[idx].CreatorID

			creatorName := getCreator(creatorID)

			// strMember := getMember(idProject)

			// if strMember == "unknown" {
			// 	strMember = getCreator(creatorID)
			// }
			arrStr = append(arrStr, strconv.Itoa(idProject))
			arrStr = append(arrStr, "|")
			arrStr = append(arrStr, nameProject)
			arrStr = append(arrStr, "|")
			// arrStr = append(arrStr, descProject)
			// arrStr = append(arrStr, "|")
			arrStr = append(arrStr, urlProject)
			arrStr = append(arrStr, "|")
			// arrStr = append(arrStr, strMember)
			arrStr = append(arrStr, onlyDate)
			arrStr = append(arrStr, "|")
			arrStr = append(arrStr, creatorName)
			arrStr = append(arrStr, "\n")

			idx++
			a++
		}
		i++
	}

	strWrite := strings.Join(arrStr, "")

	filename := "allRepo.csv"

	tes, errWri := writeToFile(filename, strWrite)

	if errWri != nil {
		log.Fatalln(errWri)
	}

	fmt.Println(tes)
	return a
}

func writeToFile(filename, data string) (string, error) {

	onlyName := filename

	file, err := os.Create("./gitlab-project/" + onlyName)
	if err != nil {
		fmt.Println(err, "err21")
		return "", err
	}
	defer file.Close()
	_, err = io.WriteString(file, data)
	if err != nil {
		return "", err
	}
	return onlyName, file.Sync()
}

func getSubGroup(email string) (error, string) {
	var connMysq *sql.DB

	dbUrl := "@tcp(10.1.35.15:3306)/smartfren?charset=utf8"

	var err error

	connMysq, err = sql.Open("mysql", dbUrl)

	defer connMysq.Close()

	if err != nil {
		log.Println(err, "cek")
		return err, ""
	}

	stmt := `select sub_group from sap_data_pentaho sdp where email = '` + email + `';`

	activeRecord, errQuery := connMysq.Query(stmt)

	defer activeRecord.Close()

	if errQuery != nil {
		log.Println("Error", errQuery)
		return errQuery, ""
	}

	subGroup := ""
	for activeRecord.Next() {
		err = activeRecord.Scan(&subGroup)
		// resp = append(resp,bk)
	}

	if err != nil {
		log.Println("Error", err)
		return err, ""
	}

	fmt.Println(subGroup)

	return nil, subGroup

}

func readDb(email string) (error, *DataKaryawan) {
	// var resp []*DataKaryawan
	resp := new(DataKaryawan)

	var err error

	var connMysq *sql.DB

	dbUrl := "@tcp(10.1.35.15:3306)/smartfren?charset=utf8"

	connMysq, err = sql.Open("mysql", dbUrl)

	defer connMysq.Close()

	if err != nil {
		log.Println(err, "cek")
		return err, resp
	}

	stmt := `select nik,name,join_date,superior_name,superior_email from sap_data_pentaho sdp where email = '` + email + `';`

	activeRecord, errQuery := connMysq.Query(stmt)

	defer activeRecord.Close()

	if errQuery != nil {
		log.Println("Error", errQuery)
		return errQuery, resp
	}

	for activeRecord.Next() {
		err = activeRecord.Scan(&resp.Nik, &resp.Nama, &resp.Join, &resp.SuperiorName, &resp.SuperiorEmail)
		// resp = append(resp,bk)
	}

	if err != nil {
		log.Println("Error", err)
		return err, resp
	}

	return nil, resp

}

func getKaryawan() (error, []*DataKaryawan) {
	var resp []*DataKaryawan

	var err error

	var connMysq *sql.DB

	dbUrl := "@tcp(10.1.35.15:3306)/smartfren?charset=utf8"

	connMysq, err = sql.Open("mysql", dbUrl)

	defer connMysq.Close()

	if err != nil {
		log.Println(err, "cek")
		return err, resp
	}

	stmt := `select nik,name,email,join_date,superior_name,superior_email,groups from sap_data_pentaho sdp;`

	activeRecord, errQuery := connMysq.Query(stmt)

	defer activeRecord.Close()

	if errQuery != nil {
		log.Println("Error", errQuery)
		return errQuery, resp
	}

	for activeRecord.Next() {
		bk := new(DataKaryawan)
		err = activeRecord.Scan(&bk.Nik, &bk.Nama, &bk.Email, &bk.Join, &bk.SuperiorName, &bk.SuperiorEmail, &bk.Groups)
		resp = append(resp, bk)
	}

	if err != nil {
		log.Println("Error", err)
		return err, resp
	}

	return nil, resp

}

func getSuperior(email string, allKaryawan []*DataKaryawan) (string, string) {
	newEmail := strings.ToUpper(email)

	superior1 := ""
	superior2 := ""

	for _, rangeAll := range allKaryawan {
		if newEmail == rangeAll.Email {
			superior1 = rangeAll.SuperiorEmail
			superior2 = getSuperior2(superior1, allKaryawan)
		}
	}

	return superior1, superior2
}

func getSuperior2(email string, allKaryawan []*DataKaryawan) string {
	newEmail := strings.ToUpper(email)

	superior := ""

	for _, rangeAll := range allKaryawan {
		if newEmail == rangeAll.Email {
			superior = rangeAll.SuperiorEmail
		}
	}

	return superior
}

func writeToExcel(gitlabList []*GitlabList) error {

	xlsx := excelize.NewFile()

	style, errStyle := xlsx.NewStyle(`{
	    "font": {
	        "bold": true,
	        "size": 15
	    },
	    "fill": {
	        "type": "pattern",
	        "color": ["#E0EBF5"],
	        "pattern": 1
	    }
	}`)

	if errStyle != nil {
		log.Println(errStyle)
	}

	sheet1Name := "detail-repo"
	xlsx.SetSheetName(xlsx.GetSheetName(1), sheet1Name)

	xlsx.SetCellStyle(sheet1Name, "A2", "G2", style)

	xlsx.SetColWidth(sheet1Name, "D", "D", 21)
	xlsx.SetColWidth(sheet1Name, "F", "G", 11)

	xlsx.SetCellValue(sheet1Name, "A2", "No")
	xlsx.SetCellValue(sheet1Name, "B2", "Project Name")
	xlsx.SetCellValue(sheet1Name, "C2", "URL Project")
	xlsx.SetCellValue(sheet1Name, "D2", "Owner")
	xlsx.SetCellValue(sheet1Name, "E2", "Member")
	xlsx.SetCellValue(sheet1Name, "F2", "Superior 1")
	xlsx.SetCellValue(sheet1Name, "G2", "Superior 2")

	err := xlsx.AutoFilter(sheet1Name, "A2", "C2", "")
	if err != nil {
		log.Fatal("ERROR", err.Error())
		return err
	}

	for i, rangeGitlab := range gitlabList {
		xlsx.SetCellValue(sheet1Name, fmt.Sprintf("A%d", i+3), strconv.Itoa(i+1))
		xlsx.SetCellValue(sheet1Name, fmt.Sprintf("B%d", i+3), rangeGitlab.NameProject)
		xlsx.SetCellValue(sheet1Name, fmt.Sprintf("C%d", i+3), rangeGitlab.UrlProject)
		xlsx.SetCellValue(sheet1Name, fmt.Sprintf("D%d", i+3), rangeGitlab.Owner)
		xlsx.SetCellValue(sheet1Name, fmt.Sprintf("E%d", i+3), rangeGitlab.Member)
		xlsx.SetCellValue(sheet1Name, fmt.Sprintf("F%d", i+3), rangeGitlab.Superior1)
		xlsx.SetCellValue(sheet1Name, fmt.Sprintf("G%d", i+3), rangeGitlab.Superior2)
	}

	err = xlsx.SaveAs("/var/gitlab-project/gitlab-project.xlsx")
	if err != nil {
		fmt.Println(err)
		return err
	}
	return nil
}
