package main


type Client struct {
  inDB bool
  Name string
  Email string
  Phone string
}

func createClient(name string, email string, phone string) *Client{
  return &Client{false, name, email, phone}
}


func main() {
  
  s := Sql{}
  s.Init()
  defer s.Close()

  s.CreateTables()

  eduardo := createClient("Eduardo", "eduardo.schulz@example.com", "555-5555")
  s.ClientInsertion(*eduardo)

}


