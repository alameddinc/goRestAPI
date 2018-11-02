package main

import (
    "encoding/json"
    "log"
    "net/http"
    "strconv"
    "fmt"
    "github.com/gorilla/mux"
)
type Skill struct{
  Name string
  Level int64
}

type Person struct{
  ID int64
  FirstName string
  Surname string
  Phone string
  Age int64
  Email string
  City string
  Skills []Skill
}

type Task struct{
  ID int64
  title string
  description string
  userId int64
}

var people []Person

// our main function
func main() {
    var mySkill []Skill
    php := Skill{Name:"PHP",Level:76}
    js := Skill{Name:"Javascript",Level:87}
    mySkill = append(mySkill,php)
    mySkill = append(mySkill,js)
    people = append(people,Person{ID:1,FirstName:"Alameddin",Surname:"Çelik",Phone:"5417907817",Email:"alameddinc@gmail.com",City:"Istanbul",Age:27,Skills:mySkill})
    router := mux.NewRouter()
    router.HandleFunc("/all", GetPeople).Methods("GET")
    router.HandleFunc("/add", AddDataWithGet).Methods("GET")
    router.HandleFunc("/delete", DeleteData).Methods("GET")
    router.HandleFunc("/update", UpdateData).Methods("GET")
    router.HandleFunc("/find", FindData).Methods("GET")
    router.HandleFunc("/filter", FilterAgeData).Methods("GET")
    log.Fatal(http.ListenAndServe(":8000", router))
}
//All Contacts
func GetPeople(w http.ResponseWriter, r *http.Request)  {
  json.NewEncoder(w).Encode(people)
}
//Add Contact Get and post
func AddDataWithGet(w http.ResponseWriter, r *http.Request){
  params := r.URL.Query()
  message := "failed"
  if (params["id"]!=nil && params["firstname"]!=nil && params["surname"]!=nil && params["phone"]!=nil && params["age"] != nil){
    tempAge, _ := strconv.ParseInt(params["age"][0], 10, 64)
    tempId, _ := strconv.ParseInt(params["id"][0], 10, 64)
    firstname :=params["firstname"][0]
    surname := params["surname"][0]
    phone := params["phone"][0]
    if(!isSet(tempId)){
      message = "Success"
      people = append(people, Person{ID:tempId,FirstName:firstname,Surname:surname,Phone:phone,Age:tempAge})
    }
  }
  json.NewEncoder(w).Encode(message)
}
//Delete Contact
func DeleteData(w http.ResponseWriter, r *http.Request){
  id := r.URL.Query()["id"][0]
  tempId , _ := strconv.ParseInt(id, 10, 64)
  var temp []Person
  for _, p := range people {
    if p.ID != tempId{
      temp = append(temp,p)
    }
    people = temp
  }
}
//Update Contacts
func UpdateData(w http.ResponseWriter,r *http.Request){
  count := 0
  id , _ := strconv.ParseInt(r.URL.Query()["id"][0],10,64)
  var temp []Person
  for _, p:= range people {
    if(p.ID == id){
        count++
        if(r.URL.Query()["firstname"]!=nil){
          p.FirstName=r.URL.Query()["firstname"][0]
        }
        if(r.URL.Query()["surname"]!=nil){
          p.Surname=r.URL.Query()["surname"][0]
        }
        if(r.URL.Query()["age"]!=nil){
          p.Age,_=strconv.ParseInt(r.URL.Query()["age"][0],10,64)
        }
        if(r.URL.Query()["phone"]!=nil){
          p.Phone=r.URL.Query()["phone"][0]
        }
    }
    temp = append(temp, p)
  }
  people = temp
  //Sprintf örneği içermesi için bu yöntemi kullanmış bulunmaktayım.
  json.NewEncoder(w).Encode(fmt.Sprintf("%d adet veri güncellendi",count))
}

//Find Contact
func FindData(w http.ResponseWriter, r *http.Request){
  status := false
  id,_ := strconv.ParseInt(r.URL.Query()["id"][0],10,64)
  for _,p := range people {
    if(p.ID == id){
      json.NewEncoder(w).Encode(p)
      status = true
    }
  }
  if(!status){
    json.NewEncoder(w).Encode("NOT FOUND")
  }
}
//filterWithAge
func FilterAgeData(w http.ResponseWriter, r *http.Request){
  count := 0
  var temp []Person
  min,_ := strconv.ParseInt(r.URL.Query()["min"][0],10,64)
  max,_ := strconv.ParseInt(r.URL.Query()["max"][0],10,64)
  for _,p := range people {
    if(p.Age >= min && p.Age <= max){
      temp = append(temp,p)
      count++
    }
  }
  if(count==0){
    json.NewEncoder(w).Encode("NOT FOUND")
  }else{
    json.NewEncoder(w).Encode(temp)
  }
}

//is Set
func isSet(id int64) bool{
  for _,p := range people {
    if p.ID == id{
      return true
    }
  }
  return false
}
