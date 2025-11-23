package img

import (
	"unsafe"

	"github.com/SunshineZzzz/purego-sdl3/sdl"
)

type Animation struct {
	W, H   int32
	Count  int32
	frames **sdl.Surface
	delays *int32
}

func (a *Animation) Frames() []*sdl.Surface {
	return unsafe.Slice(a.frames, a.Count)
}

func (a *Animation) Delays() []int32 {
	return unsafe.Slice(a.delays, a.Count)
}

// Version gets the version of the dynamically linked SDL_image library.
func Version() (major, minor, patch int32) {
	version := imgVersion()
	major = version / 1000000
	minor = version / 1000 % 1000
	patch = version % 1000
	return
}

// FreeAnimation disposes an Animation and frees its resources.
func FreeAnimation(anim *Animation) {
	imgFreeAnimation(anim)
}

func IsAVIF(src *sdl.IOStream) bool {
	return imgIsAVIF(src)
}

func IsBMP(src *sdl.IOStream) bool {
	return imgIsBMP(src)
}

func IsCUR(src *sdl.IOStream) bool {
	return imgIsCUR(src)
}

func IsGIF(src *sdl.IOStream) bool {
	return imgIsGIF(src)
}

func IsICO(src *sdl.IOStream) bool {
	return imgIsICO(src)
}

func IsJPG(src *sdl.IOStream) bool {
	return imgIsJPG(src)
}

func IsJXL(src *sdl.IOStream) bool {
	return imgIsJXL(src)
}

func IsLBM(src *sdl.IOStream) bool {
	return imgIsLBM(src)
}

func IsPCX(src *sdl.IOStream) bool {
	return imgIsPCX(src)
}

func IsPNG(src *sdl.IOStream) bool {
	return imgIsPNG(src)
}

func IsPNM(src *sdl.IOStream) bool {
	return imgIsPNM(src)
}

func IsQOI(src *sdl.IOStream) bool {
	return imgIsQOI(src)
}

func IsSVG(src *sdl.IOStream) bool {
	return imgIsSVG(src)
}

func IsTIF(src *sdl.IOStream) bool {
	return imgIsTIF(src)
}

func IsWEBP(src *sdl.IOStream) bool {
	return imgIsWEBP(src)
}

func IsXCF(src *sdl.IOStream) bool {
	return imgIsXCF(src)
}

func IsXPM(src *sdl.IOStream) bool {
	return imgIsXPM(src)
}

func IsXV(src *sdl.IOStream) bool {
	return imgIsXV(src)
}

func Load(file string) *sdl.Surface {
	return imgLoad(file)
}

func LoadIO(src *sdl.IOStream, closeio bool) *sdl.Surface {
	return imgLoadIO(src, closeio)
}

func LoadAnimation(file string) *Animation {
	return imgLoadAnimation(file)
}

func LoadAnimationIO(src *sdl.IOStream, closeio bool) *Animation {
	return imgLoadAnimationIO(src, closeio)
}

func LoadAnimationTypedIO(src *sdl.IOStream, closeio bool, format string) *Animation {
	return imgLoadAnimationTypedIO(src, closeio, format)
}

func LoadAVIFIO(src *sdl.IOStream) *sdl.Surface {
	return imgLoadAVIFIO(src)
}

func LoadBMPIO(src *sdl.IOStream) *sdl.Surface {
	return imgLoadBMPIO(src)
}

func LoadCURIO(src *sdl.IOStream) *sdl.Surface {
	return imgLoadCURIO(src)
}

func LoadGIFIO(src *sdl.IOStream) *sdl.Surface {
	return imgLoadGIFIO(src)
}

func LoadGIFAnimationIO(src *sdl.IOStream) *Animation {
	return imgLoadGIFAnimationIO(src)
}

func LoadICOIO(src *sdl.IOStream) *sdl.Surface {
	return imgLoadICOIO(src)
}

func LoadJPGIO(src *sdl.IOStream) *sdl.Surface {
	return imgLoadJPGIO(src)
}

func LoadJXLIO(src *sdl.IOStream) *sdl.Surface {
	return imgLoadJXLIO(src)
}

func LoadLBMIO(src *sdl.IOStream) *sdl.Surface {
	return imgLoadLBMIO(src)
}

func LoadPCXIO(src *sdl.IOStream) *sdl.Surface {
	return imgLoadPCXIO(src)
}

func LoadPNGIO(src *sdl.IOStream) *sdl.Surface {
	return imgLoadPNGIO(src)
}

func LoadPNMIO(src *sdl.IOStream) *sdl.Surface {
	return imgLoadPNMIO(src)
}

func LoadQOIIO(src *sdl.IOStream) *sdl.Surface {
	return imgLoadQOIIO(src)
}

func LoadSizedSVGIO(src *sdl.IOStream, width int32, height int32) *sdl.Surface {
	return imgLoadSizedSVGIO(src, width, height)
}

func LoadSVGIO(src *sdl.IOStream) *sdl.Surface {
	return imgLoadSVGIO(src)
}

func LoadTexture(renderer *sdl.Renderer, file string) *sdl.Texture {
	return imgLoadTexture(renderer, file)
}

func LoadTextureIO(renderer *sdl.Renderer, src *sdl.IOStream, closeio bool) *sdl.Texture {
	return imgLoadTextureIO(renderer, src, closeio)
}

func LoadTextureTypedIO(renderer *sdl.Renderer, src *sdl.IOStream, closeio bool, format string) *sdl.Texture {
	return imgLoadTextureTypedIO(renderer, src, closeio, format)
}

func LoadTGAIO(src *sdl.IOStream) *sdl.Surface {
	return imgLoadTGAIO(src)
}

func LoadTIFIO(src *sdl.IOStream) *sdl.Surface {
	return imgLoadTIFIO(src)
}

func LoadTypedIO(src *sdl.IOStream, closeio bool, format string) *sdl.Surface {
	return imgLoadTypedIO(src, closeio, format)
}

func LoadWEBPIO(src *sdl.IOStream) *sdl.Surface {
	return imgLoadWEBPIO(src)
}

func LoadWEBPAnimationIO(src *sdl.IOStream) *Animation {
	return imgLoadWEBPAnimationIO(src)
}

func LoadXCFIO(src *sdl.IOStream) *sdl.Surface {
	return imgLoadXCFIO(src)
}

func LoadXPMIO(src *sdl.IOStream) *sdl.Surface {
	return imgLoadXPMIO(src)
}

func LoadXVIO(src *sdl.IOStream) *sdl.Surface {
	return imgLoadXVIO(src)
}

func ReadXPMFromArray(xpm **byte) *sdl.Surface {
	return imgReadXPMFromArray(xpm)
}

func ReadXPMFromArrayToRGB888(xpm **byte) *sdl.Surface {
	return imgReadXPMFromArrayToRGB888(xpm)
}

func SaveAVIF(surface *sdl.Surface, file string, quality int32) bool {
	return imgSaveAVIF(surface, file, quality)
}

func SaveAVIFIO(surface *sdl.Surface, dst *sdl.IOStream, closeio bool, quality int32) bool {
	return imgSaveAVIFIO(surface, dst, closeio, quality)
}

func SaveJPG(surface *sdl.Surface, file string, quality int32) bool {
	return imgSaveJPG(surface, file, quality)
}

func SaveJPGIO(surface *sdl.Surface, dst *sdl.IOStream, closeio bool, quality int32) bool {
	return imgSaveJPGIO(surface, dst, closeio, quality)
}

func SavePNG(surface *sdl.Surface, file string) bool {
	return imgSavePNG(surface, file)
}

func SavePNGIO(surface *sdl.Surface, dst *sdl.IOStream, closeio bool) bool {
	return imgSavePNGIO(surface, dst, closeio)
}
