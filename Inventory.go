package main
type Inventory struct{
	Products []*Product
	Availability map[string]int

}
func NewInventory()*Inventory {
	return &Inventory{}
}
func(i *Inventory) AddProduct(P *Product){
	if i.Availability==nil{
		i.Availability=make(map[string]int)
	}
	if i.Availability[P.Name]==0{
		i.Products=append(i.Products,P)
		i.Availability[P.Name]++
	}else{
		i.Availability[P.Name]++
	}
	
	
}
func(i *Inventory) RemoveProduct(P *Product){
	var ret Inventory
	for _,curr:=range i.Products{
		if curr!=P{
			ret.Products=append(ret.Products,curr)
		}
	}
	if i.Availability[P.Name]!=0{
	i.Availability[P.Name]--
	}
}

func(i *Inventory) Quantity (P *Product) int{
  m:=i.Availability[P.Name]
  return m
} 
