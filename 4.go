type Girl struct {

}

type GroupLeader struct {
	girls []Girl
}

type Teacher struct{}


func (g GroupLeader) CountGrils()  {
	fmt.Println("the sum of girls is ", len(g.girls))
}

func (t Teacher) Command(leader GroupLeader)  {
	
}