package Beer

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

const URLbeer = "https://api.punkapi.com/v2/beers"

type Volume struct {
	Value float32 `json:"value"`
	Unit  string  `json:"unit"`
}
type formulaBeer struct {
	Id          int     `json:"id"`
	Name        string  `json:"name"`
	Tagline     string  `json:"tagline"`
	Description string  `json:"description"`
	Ebc         float32 `json:"ebc"`
	Ibu         float32 `json:"ibu"`
	Volume      Volume  `json:"volume"`
}

func (Beer *formulaBeer) Info() {
	str := "______________________\n" +
		" name: " + Beer.Name +
		"\n Tag line: " + Beer.Tagline +
		"\n description: " + Beer.Description +
		"\n ebc: " + fmt.Sprint(Beer.Ebc) +
		"\n ibu: " + fmt.Sprint(Beer.Ibu) +
		"\n volume: " + fmt.Sprint(Beer.Volume.Value) + " " + Beer.Volume.Unit

	fmt.Println(str)

}

func GetBeers(minIBU int, maxIBU int, recipesCount int) (*[]formulaBeer, error) {
	var arrayFormulaBeer []formulaBeer

	//resp, err := http.DefaultClient.Get(URLbeer + "?ibu_gt=" + strconv.Itoa(minIBU) + "&ibu_lt=" + strconv.Itoa(maxIBU)+"&per_page="+strconv.Itoa(recipesCount))
	resp, err := http.DefaultClient.Get(fmt.Sprintf("%s?ibu_gt=%d&ibu_lt=%d&per_page=%d", URLbeer, minIBU, maxIBU, recipesCount))
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)

	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(body, &arrayFormulaBeer)
	if err != nil {
		return nil, err

	}

	return &arrayFormulaBeer, nil

}
