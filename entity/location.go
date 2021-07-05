package entity
type Location struct{
    Longitude float32
	Latitude float32 
	city string
   
   }
func (cit Location) GetCity() string{
	   
		return cit.city
	  }