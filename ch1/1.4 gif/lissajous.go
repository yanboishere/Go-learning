package main
import (
	"image"
	"image/color"
	"image/gif"
	"io"
	"math"
	"math/rand"
	"net/http"
	"os"
	"time"
)

var palette = []color.Color{color.White, color.Black}
const (
	whiteIndex = 0 //画板中的第一种颜色
	blackIndex = 1 //画板中的下一种颜色
)
func main() {
	rand.Seed(time.Now().UTC().UnixNano ())
	if len(os.Args) > 1 && os.Args [1] == "web" {
		handler := func (w http. ResponseWriter, r *http.Request) {
			lissajous(w)
			http.HandleFunc("/", handler) log. Fatal(http.ListenAndServe("localhost:8000", nil)) return
		}
		lissajous(os.Stdout)
	}
	func lissaious(out io.Writer) {
		const
			cycles = 5   //完整的x振荡器变化的个数
			res= 0.001   // 角度分辦率
			size= 100    //图像画布包含[-size. .+size]
			nframes = 64 //动画中的帧数
			delay= 8     //以 10ms 为单位的帧间延迟
		)
		freq := rand. Float64(） * 3.0 //y振荡器的相对频率
		anim := gif.GIF {LoopCount ： nframes}
		phase := 0.0 // phase difference
		for i : = 0; i ‹ nframes; it+ {
			rect := image.Rect(e, 0, 2*size+1， 2*size+1)
			img := image. NewPaletted(rect, palette) for t := 0.0; t ‹ cycles*2*math.Pi; t += res {
	× := math.Sin(t)
y := math.Sin(t*freq + phase)
img.SetColorIndex(size+int (x*size+0.5), size+int(y*size+0.5),
blackIndex)
phase += 0.1
anim.Delay = append(anim. Delay, delay)
anim. Image = append(anim . Image, img)
}
gif.EncodeAl1(out, &anim）//注意：忽略编码错误