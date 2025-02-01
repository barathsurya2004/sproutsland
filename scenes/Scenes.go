package scenes

type Scenes interface {
	Update() int
	OnFirstLoad()
}
