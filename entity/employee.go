package entity 


type Employee struct{
 Name string 
 Age int
 Location Location
 Gender string

}
func (emp Employee) GetCity() string{
	
      return emp.Location.city
   }