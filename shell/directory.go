package shell

type directory struct {
	name string
	desc string
	parent *directory
	child []*directory
	action action
}

func (d *directory) add(child *directory){
	d.child = append(d.child, child)
	child.parent = d
}

func (d *directory) reset(){
	d.child = make([]*directory, 0)
}

func newDirectory(name string, desc string, actions ...actions) *directory{
	return &directory{
		name: name,
		desc: desc,
		action: func(s *shell, inputs []byte) (string, error) {
			return distribute(s, inputs, actions...)
		},
	}
}

var root = newDirectory("/", "", rootActions, baseActions)

func init(){
}
