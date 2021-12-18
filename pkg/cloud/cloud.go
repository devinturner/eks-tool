package cloud

type Cloud struct {
	Regions []*Region `json:"regions,omitempty"`
}
