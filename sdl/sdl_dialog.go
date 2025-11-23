package sdl

import (
	"unsafe"

	"github.com/ebitengine/purego"
	"github.com/SunshineZzzz/purego-sdl3/internal/convert"
)

const (
	PropFileDialogFiltersPointer = "SDL.filedialog.filters"
	PropFileDialogNfiltersNumber = "SDL.filedialog.nfilters"
	PropFileDialogWindowPointer  = "SDL.filedialog.window"
	PropFileDialogLocationString = "SDL.filedialog.location"
	PropFileDialogManyBoolean    = "SDL.filedialog.many"
	PropFileDialogTitleString    = "SDL.filedialog.title"
	PropFileDialogAcceptString   = "SDL.filedialog.accept"
	PropFileDialogCancelString   = "SDL.filedialog.cancel"
)

type FileDialogType uint32

const (
	FileDialogOpenFile FileDialogType = iota
	FileDialogSaveFile
	FileDialogOpenFolder
)

type DialogFileFilter struct {
	name, pattern *byte
}

func NewDialogFileFilter(name, pattern string) DialogFileFilter {
	return DialogFileFilter{convert.ToBytePtr(name), convert.ToBytePtr(pattern)}
}

func (d DialogFileFilter) Name() string {
	return convert.ToString(d.name)
}

func (d DialogFileFilter) Pattern() string {
	return convert.ToString(d.pattern)
}

type DialogFileCallback uintptr

func NewDialogFileCallback(callback func(userdata unsafe.Pointer, filelist []string, filter int32)) DialogFileCallback {
	cb := purego.NewCallback(func(userdata unsafe.Pointer, filelist **byte, filter int32) uintptr {
		callback(userdata, convert.ToStringSlice(filelist), filter)
		return 0
	})

	return DialogFileCallback(cb)
}

func ShowFileDialogWithProperties(dialogType FileDialogType, callback DialogFileCallback, userdata unsafe.Pointer, props PropertiesID) {
	sdlShowFileDialogWithProperties(dialogType, callback, userdata, props)
}

func ShowOpenFileDialog(callback DialogFileCallback, userdata unsafe.Pointer, window *Window, filters []DialogFileFilter, defaultLocation string, allowMany bool) {
	sdlShowOpenFileDialog(callback, userdata, window, filters, int32(len(filters)), convert.ToBytePtrNullable(defaultLocation), allowMany)
}

func ShowOpenFolderDialog(callback DialogFileCallback, userdata unsafe.Pointer, window *Window, defaultLocation string, allowMany bool) {
	sdlShowOpenFolderDialog(callback, userdata, window, defaultLocation, allowMany)
}

func ShowSaveFileDialog(callback DialogFileCallback, userdata unsafe.Pointer, window *Window, filters []DialogFileFilter, defaultLocation string) {
	sdlShowSaveFileDialog(callback, userdata, window, filters, int32(len(filters)), defaultLocation)
}
