package alphabeta

type BattlesnakeInfoResponse struct {
	APIVersion string `json:"apiversion"`
	Author     string `json:"author"`
	Color      string `json:"color"`
	Head       string `json:"head"`
	Tail       string `json:"tail"`
}
type MoveResponse struct {
	Move  string `json:"move"`
	Shout string `json:"shout,omitempty"`
}