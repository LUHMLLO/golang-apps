package main

type Getter interface {
	Get()
}
type Putter interface {
	Put()
}
type Deleter interface {
	Delete()
}
type Patcher interface {
	Patch()
}

type Storer interface {
	Getter
	Putter
}
