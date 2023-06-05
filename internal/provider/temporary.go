package provider

type Temporary struct {
	BaseDir string
}

func NewTemporary(baseDir string) *Temporary {
	return &Temporary{
		BaseDir: baseDir,
	}
}
