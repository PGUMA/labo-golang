package interace

type Wizard interface {
	magic()
}

type マッシュバーンデッド struct{}

type レインエイムズ struct{}

func (w *レインエイムズ) magic() {
	println("パルチザン")
}

var _ Wizard = (*レインエイムズ)(nil)

// var _ Wizard = (*マッシュバーンデッド)(nil) This cmopile error
