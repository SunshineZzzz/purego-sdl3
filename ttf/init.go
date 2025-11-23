package ttf

import (
	"runtime"

	"github.com/SunshineZzzz/purego-sdl3/internal/shared"
	"github.com/SunshineZzzz/purego-sdl3/sdl"
	"github.com/ebitengine/purego"
)

var (
	ttfAddFallbackFont                        func(*Font, *Font) bool
	ttfAppendTextString                       func(*Text, string, uint64) bool
	ttfClearFallbackFonts                     func(*Font)
	ttfCloseFont                              func(*Font)
	ttfCopyFont                               func(*Font) *Font
	ttfCreateGPUTextEngine                    func(*sdl.GPUDevice) *TextEngine
	ttfCreateGPUTextEngineWithProperties      func(sdl.PropertiesID) *TextEngine
	ttfCreateRendererTextEngine               func(*sdl.Renderer) *TextEngine
	ttfCreateRendererTextEngineWithProperties func(sdl.PropertiesID) *TextEngine
	ttfCreateSurfaceTextEngine                func() *TextEngine
	ttfCreateText                             func(*TextEngine, *Font, string, uint64) *Text
	ttfDeleteTextString                       func(*Text, int32, int32) bool
	ttfDestroyGPUTextEngine                   func(*TextEngine)
	ttfDestroyRendererTextEngine              func(*TextEngine)
	ttfDestroySurfaceTextEngine               func(*TextEngine)
	ttfDestroyText                            func(*Text)
	ttfDrawRendererText                       func(*Text, float32, float32) bool
	ttfDrawSurfaceText                        func(*Text, int32, int32, *sdl.Surface) bool
	ttfFontHasGlyph                           func(*Font, rune) bool
	ttfFontIsFixedWidth                       func(*Font) bool
	ttfFontIsScalable                         func(*Font) bool
	ttfGetFontAscent                          func(*Font) int32
	ttfGetFontDescent                         func(*Font) int32
	ttfGetFontDirection                       func(*Font) Direction
	ttfGetFontDPI                             func(*Font, *int32, *int32) bool
	ttfGetFontFamilyName                      func(*Font) string
	ttfGetFontGeneration                      func(*Font) uint32
	ttfGetFontHeight                          func(*Font) int32
	ttfGetFontHinting                         func(*Font) HintingFlags
	ttfGetFontKerning                         func(*Font) bool
	ttfGetFontLineSkip                        func(*Font) int32
	ttfGetFontOutline                         func(*Font) int32
	ttfGetFontProperties                      func(*Font) sdl.PropertiesID
	ttfGetFontScript                          func(*Font) uint32
	ttfGetFontSDF                             func(*Font) bool
	ttfGetFontSize                            func(*Font) float32
	ttfGetFontStyle                           func(*Font) FontStyleFlags
	ttfGetFontStyleName                       func(*Font) string
	ttfGetFontWrapAlignment                   func(*Font) HorizontalAlignment
	ttfGetFreeTypeVersion                     func(*int32, *int32, *int32)
	ttfGetGlyphImage                          func(*Font, rune, *ImageType) *sdl.Surface
	ttfGetGlyphImageForIndex                  func(*Font, uint32, *ImageType) *sdl.Surface
	ttfGetGlyphKerning                        func(*Font, rune, rune, *int32) bool
	ttfGetGlyphMetrics                        func(*Font, rune, *int32, *int32, *int32, *int32, *int32) bool
	ttfGetGlyphScript                         func(rune) uint32
	ttfGetGPUTextDrawData                     func(*Text) *GPUAtlasDrawSequence
	ttfGetGPUTextEngineWinding                func(*TextEngine) GPUTextEngineWinding
	ttfGetHarfBuzzVersion                     func(*int32, *int32, *int32)
	ttfGetNextTextSubString                   func(*Text, *SubString, *SubString) bool
	ttfGetNumFontFaces                        func(*Font) int32
	ttfGetPreviousTextSubString               func(*Text, *SubString, *SubString) bool
	ttfGetStringSize                          func(*Font, string, uint64, *int32, *int32) bool
	ttfGetStringSizeWrapped                   func(*Font, string, uint64, int32, *int32, *int32) bool
	ttfGetTextColor                           func(*Text, *uint8, *uint8, *uint8, *uint8) bool
	ttfGetTextColorFloat                      func(*Text, *float32, *float32, *float32, *float32) bool
	ttfGetTextDirection                       func(*Text) Direction
	ttfGetTextEngine                          func(*Text) *TextEngine
	ttfGetTextFont                            func(*Text) *Font
	ttfGetTextPosition                        func(*Text, *int32, *int32) bool
	ttfGetTextProperties                      func(*Text) sdl.PropertiesID
	ttfGetTextScript                          func(*Text) uint32
	ttfGetTextSize                            func(*Text, *int32, *int32) bool
	ttfGetTextSubString                       func(*Text, int32, *SubString) bool
	ttfGetTextSubStringForLine                func(*Text, int32, *SubString) bool
	ttfGetTextSubStringForPoint               func(*Text, int32, int32, *SubString) bool
	ttfGetTextSubStringsForRange              func(*Text, int32, int32, *int32)
	ttfGetTextWrapWidth                       func(*Text, *int32) bool
	ttfInit                                   func() bool
	ttfInsertTextString                       func(*Text, int32, string, uint64) bool
	ttfMeasureString                          func(*Font, string, uint64, int32, *int32, *uint64) bool
	ttfOpenFont                               func(string, float32) *Font
	ttfOpenFontIO                             func(*sdl.IOStream, bool, float32) *Font
	ttfOpenFontWithProperties                 func(sdl.PropertiesID) *Font
	ttfQuit                                   func()
	ttfRemoveFallbackFont                     func(*Font, *Font)
	ttfRenderGlyphBlended                     func(*Font, rune, uintptr) *sdl.Surface
	ttfRenderGlyphLCD                         func(*Font, rune, uintptr, uintptr) *sdl.Surface
	ttfRenderGlyphShaded                      func(*Font, rune, uintptr, uintptr) *sdl.Surface
	ttfRenderGlyphSolid                       func(*Font, rune, uintptr) *sdl.Surface
	ttfRenderTextBlended                      func(*Font, string, uint64, uintptr) *sdl.Surface
	ttfRenderTextBlendedWrapped               func(*Font, string, uint64, uintptr, int32) *sdl.Surface
	ttfRenderTextLCD                          func(*Font, string, uint64, uintptr, uintptr) *sdl.Surface
	ttfRenderTextLCDWrapped                   func(*Font, string, uint64, uintptr, uintptr, int32) *sdl.Surface
	ttfRenderTextShaded                       func(*Font, string, uint64, uintptr, uintptr) *sdl.Surface
	ttfRenderTextShadedWrapped                func(*Font, string, uint64, uintptr, uintptr, int32) *sdl.Surface
	ttfRenderTextSolid                        func(*Font, string, uint64, uintptr) *sdl.Surface
	ttfRenderTextSolidWrapped                 func(*Font, string, uint64, uintptr, int32) *sdl.Surface
	ttfSetFontDirection                       func(*Font, Direction) bool
	ttfSetFontHinting                         func(*Font, HintingFlags)
	ttfSetFontKerning                         func(*Font, bool)
	ttfSetFontLanguage                        func(*Font, string) bool
	ttfSetFontLineSkip                        func(*Font, int32)
	ttfSetFontOutline                         func(*Font, int32) bool
	ttfSetFontScript                          func(*Font, uint32) bool
	ttfSetFontSDF                             func(*Font, bool) bool
	ttfSetFontSize                            func(*Font, float32) bool
	ttfSetFontSizeDPI                         func(*Font, float32, int32, int32) bool
	ttfSetFontStyle                           func(*Font, FontStyleFlags)
	ttfSetFontWrapAlignment                   func(*Font, HorizontalAlignment)
	ttfSetGPUTextEngineWinding                func(*TextEngine, GPUTextEngineWinding)
	ttfSetTextColor                           func(*Text, uint8, uint8, uint8, uint8) bool
	ttfSetTextColorFloat                      func(*Text, float32, float32, float32, float32) bool
	ttfSetTextDirection                       func(*Text, Direction) bool
	ttfSetTextEngine                          func(*Text, *TextEngine) bool
	ttfSetTextFont                            func(*Text, *Font) bool
	ttfSetTextPosition                        func(*Text, int32, int32) bool
	ttfSetTextScript                          func(*Text, uint32) bool
	ttfSetTextString                          func(*Text, string, uint64) bool
	ttfSetTextWrapWhitespaceVisible           func(*Text, bool) bool
	ttfSetTextWrapWidth                       func(*Text, int32) bool
	ttfStringToTag                            func(string) uint32
	ttfTagToString                            func(uint32, *byte, uint64)
	ttfTextWrapWhitespaceVisible              func(*Text) bool
	ttfUpdateText                             func(*Text) bool
	ttfVersion                                func() int32
	ttfWasInit                                func() int32
)

func init() {
	runtime.LockOSThread()

	var filename string
	switch runtime.GOOS {
	case "linux", "freebsd":
		filename = "libSDL3_ttf.so.0"
	case "windows":
		filename = "SDL3_ttf.dll"
	case "darwin":
		filename = "libSDL3_ttf.dylib"
	}

	lib, err := shared.Load(filename)
	if err != nil {
		panic(err)
	}

	purego.RegisterLibFunc(&ttfAddFallbackFont, lib, "TTF_AddFallbackFont")
	purego.RegisterLibFunc(&ttfAppendTextString, lib, "TTF_AppendTextString")
	purego.RegisterLibFunc(&ttfClearFallbackFonts, lib, "TTF_ClearFallbackFonts")
	purego.RegisterLibFunc(&ttfCloseFont, lib, "TTF_CloseFont")
	purego.RegisterLibFunc(&ttfCopyFont, lib, "TTF_CopyFont")
	purego.RegisterLibFunc(&ttfCreateGPUTextEngine, lib, "TTF_CreateGPUTextEngine")
	purego.RegisterLibFunc(&ttfCreateGPUTextEngineWithProperties, lib, "TTF_CreateGPUTextEngineWithProperties")
	purego.RegisterLibFunc(&ttfCreateRendererTextEngine, lib, "TTF_CreateRendererTextEngine")
	purego.RegisterLibFunc(&ttfCreateRendererTextEngineWithProperties, lib, "TTF_CreateRendererTextEngineWithProperties")
	purego.RegisterLibFunc(&ttfCreateSurfaceTextEngine, lib, "TTF_CreateSurfaceTextEngine")
	purego.RegisterLibFunc(&ttfCreateText, lib, "TTF_CreateText")
	purego.RegisterLibFunc(&ttfDeleteTextString, lib, "TTF_DeleteTextString")
	purego.RegisterLibFunc(&ttfDestroyGPUTextEngine, lib, "TTF_DestroyGPUTextEngine")
	purego.RegisterLibFunc(&ttfDestroyRendererTextEngine, lib, "TTF_DestroyRendererTextEngine")
	purego.RegisterLibFunc(&ttfDestroySurfaceTextEngine, lib, "TTF_DestroySurfaceTextEngine")
	purego.RegisterLibFunc(&ttfDestroyText, lib, "TTF_DestroyText")
	purego.RegisterLibFunc(&ttfDrawRendererText, lib, "TTF_DrawRendererText")
	purego.RegisterLibFunc(&ttfDrawSurfaceText, lib, "TTF_DrawSurfaceText")
	purego.RegisterLibFunc(&ttfFontHasGlyph, lib, "TTF_FontHasGlyph")
	purego.RegisterLibFunc(&ttfFontIsFixedWidth, lib, "TTF_FontIsFixedWidth")
	purego.RegisterLibFunc(&ttfFontIsScalable, lib, "TTF_FontIsScalable")
	purego.RegisterLibFunc(&ttfGetFontAscent, lib, "TTF_GetFontAscent")
	purego.RegisterLibFunc(&ttfGetFontDescent, lib, "TTF_GetFontDescent")
	purego.RegisterLibFunc(&ttfGetFontDirection, lib, "TTF_GetFontDirection")
	purego.RegisterLibFunc(&ttfGetFontDPI, lib, "TTF_GetFontDPI")
	purego.RegisterLibFunc(&ttfGetFontFamilyName, lib, "TTF_GetFontFamilyName")
	purego.RegisterLibFunc(&ttfGetFontGeneration, lib, "TTF_GetFontGeneration")
	purego.RegisterLibFunc(&ttfGetFontHeight, lib, "TTF_GetFontHeight")
	purego.RegisterLibFunc(&ttfGetFontHinting, lib, "TTF_GetFontHinting")
	purego.RegisterLibFunc(&ttfGetFontKerning, lib, "TTF_GetFontKerning")
	purego.RegisterLibFunc(&ttfGetFontLineSkip, lib, "TTF_GetFontLineSkip")
	purego.RegisterLibFunc(&ttfGetFontOutline, lib, "TTF_GetFontOutline")
	purego.RegisterLibFunc(&ttfGetFontProperties, lib, "TTF_GetFontProperties")
	purego.RegisterLibFunc(&ttfGetFontScript, lib, "TTF_GetFontScript")
	purego.RegisterLibFunc(&ttfGetFontSDF, lib, "TTF_GetFontSDF")
	purego.RegisterLibFunc(&ttfGetFontSize, lib, "TTF_GetFontSize")
	purego.RegisterLibFunc(&ttfGetFontStyle, lib, "TTF_GetFontStyle")
	purego.RegisterLibFunc(&ttfGetFontStyleName, lib, "TTF_GetFontStyleName")
	purego.RegisterLibFunc(&ttfGetFontWrapAlignment, lib, "TTF_GetFontWrapAlignment")
	purego.RegisterLibFunc(&ttfGetFreeTypeVersion, lib, "TTF_GetFreeTypeVersion")
	purego.RegisterLibFunc(&ttfGetGlyphImage, lib, "TTF_GetGlyphImage")
	purego.RegisterLibFunc(&ttfGetGlyphImageForIndex, lib, "TTF_GetGlyphImageForIndex")
	purego.RegisterLibFunc(&ttfGetGlyphKerning, lib, "TTF_GetGlyphKerning")
	purego.RegisterLibFunc(&ttfGetGlyphMetrics, lib, "TTF_GetGlyphMetrics")
	purego.RegisterLibFunc(&ttfGetGlyphScript, lib, "TTF_GetGlyphScript")
	purego.RegisterLibFunc(&ttfGetGPUTextDrawData, lib, "TTF_GetGPUTextDrawData")
	purego.RegisterLibFunc(&ttfGetGPUTextEngineWinding, lib, "TTF_GetGPUTextEngineWinding")
	purego.RegisterLibFunc(&ttfGetHarfBuzzVersion, lib, "TTF_GetHarfBuzzVersion")
	purego.RegisterLibFunc(&ttfGetNextTextSubString, lib, "TTF_GetNextTextSubString")
	purego.RegisterLibFunc(&ttfGetNumFontFaces, lib, "TTF_GetNumFontFaces")
	purego.RegisterLibFunc(&ttfGetPreviousTextSubString, lib, "TTF_GetPreviousTextSubString")
	purego.RegisterLibFunc(&ttfGetStringSize, lib, "TTF_GetStringSize")
	purego.RegisterLibFunc(&ttfGetStringSizeWrapped, lib, "TTF_GetStringSizeWrapped")
	purego.RegisterLibFunc(&ttfGetTextColor, lib, "TTF_GetTextColor")
	purego.RegisterLibFunc(&ttfGetTextColorFloat, lib, "TTF_GetTextColorFloat")
	purego.RegisterLibFunc(&ttfGetTextDirection, lib, "TTF_GetTextDirection")
	purego.RegisterLibFunc(&ttfGetTextEngine, lib, "TTF_GetTextEngine")
	purego.RegisterLibFunc(&ttfGetTextFont, lib, "TTF_GetTextFont")
	purego.RegisterLibFunc(&ttfGetTextPosition, lib, "TTF_GetTextPosition")
	purego.RegisterLibFunc(&ttfGetTextProperties, lib, "TTF_GetTextProperties")
	purego.RegisterLibFunc(&ttfGetTextScript, lib, "TTF_GetTextScript")
	purego.RegisterLibFunc(&ttfGetTextSize, lib, "TTF_GetTextSize")
	purego.RegisterLibFunc(&ttfGetTextSubString, lib, "TTF_GetTextSubString")
	purego.RegisterLibFunc(&ttfGetTextSubStringForLine, lib, "TTF_GetTextSubStringForLine")
	purego.RegisterLibFunc(&ttfGetTextSubStringForPoint, lib, "TTF_GetTextSubStringForPoint")
	purego.RegisterLibFunc(&ttfGetTextSubStringsForRange, lib, "TTF_GetTextSubStringsForRange")
	purego.RegisterLibFunc(&ttfGetTextWrapWidth, lib, "TTF_GetTextWrapWidth")
	purego.RegisterLibFunc(&ttfInit, lib, "TTF_Init")
	purego.RegisterLibFunc(&ttfInsertTextString, lib, "TTF_InsertTextString")
	purego.RegisterLibFunc(&ttfMeasureString, lib, "TTF_MeasureString")
	purego.RegisterLibFunc(&ttfOpenFont, lib, "TTF_OpenFont")
	purego.RegisterLibFunc(&ttfOpenFontIO, lib, "TTF_OpenFontIO")
	purego.RegisterLibFunc(&ttfOpenFontWithProperties, lib, "TTF_OpenFontWithProperties")
	purego.RegisterLibFunc(&ttfQuit, lib, "TTF_Quit")
	purego.RegisterLibFunc(&ttfRemoveFallbackFont, lib, "TTF_RemoveFallbackFont")
	purego.RegisterLibFunc(&ttfRenderGlyphBlended, lib, "TTF_RenderGlyph_Blended")
	purego.RegisterLibFunc(&ttfRenderGlyphLCD, lib, "TTF_RenderGlyph_LCD")
	purego.RegisterLibFunc(&ttfRenderGlyphShaded, lib, "TTF_RenderGlyph_Shaded")
	purego.RegisterLibFunc(&ttfRenderGlyphSolid, lib, "TTF_RenderGlyph_Solid")
	purego.RegisterLibFunc(&ttfRenderTextBlended, lib, "TTF_RenderText_Blended")
	purego.RegisterLibFunc(&ttfRenderTextBlendedWrapped, lib, "TTF_RenderText_Blended_Wrapped")
	purego.RegisterLibFunc(&ttfRenderTextLCD, lib, "TTF_RenderText_LCD")
	purego.RegisterLibFunc(&ttfRenderTextLCDWrapped, lib, "TTF_RenderText_LCD_Wrapped")
	purego.RegisterLibFunc(&ttfRenderTextShaded, lib, "TTF_RenderText_Shaded")
	purego.RegisterLibFunc(&ttfRenderTextShadedWrapped, lib, "TTF_RenderText_Shaded_Wrapped")
	purego.RegisterLibFunc(&ttfRenderTextSolid, lib, "TTF_RenderText_Solid")
	purego.RegisterLibFunc(&ttfRenderTextSolidWrapped, lib, "TTF_RenderText_Solid_Wrapped")
	purego.RegisterLibFunc(&ttfSetFontDirection, lib, "TTF_SetFontDirection")
	purego.RegisterLibFunc(&ttfSetFontHinting, lib, "TTF_SetFontHinting")
	purego.RegisterLibFunc(&ttfSetFontKerning, lib, "TTF_SetFontKerning")
	purego.RegisterLibFunc(&ttfSetFontLanguage, lib, "TTF_SetFontLanguage")
	purego.RegisterLibFunc(&ttfSetFontLineSkip, lib, "TTF_SetFontLineSkip")
	purego.RegisterLibFunc(&ttfSetFontOutline, lib, "TTF_SetFontOutline")
	purego.RegisterLibFunc(&ttfSetFontScript, lib, "TTF_SetFontScript")
	purego.RegisterLibFunc(&ttfSetFontSDF, lib, "TTF_SetFontSDF")
	purego.RegisterLibFunc(&ttfSetFontSize, lib, "TTF_SetFontSize")
	purego.RegisterLibFunc(&ttfSetFontSizeDPI, lib, "TTF_SetFontSizeDPI")
	purego.RegisterLibFunc(&ttfSetFontStyle, lib, "TTF_SetFontStyle")
	purego.RegisterLibFunc(&ttfSetFontWrapAlignment, lib, "TTF_SetFontWrapAlignment")
	purego.RegisterLibFunc(&ttfSetGPUTextEngineWinding, lib, "TTF_SetGPUTextEngineWinding")
	purego.RegisterLibFunc(&ttfSetTextColor, lib, "TTF_SetTextColor")
	purego.RegisterLibFunc(&ttfSetTextColorFloat, lib, "TTF_SetTextColorFloat")
	purego.RegisterLibFunc(&ttfSetTextDirection, lib, "TTF_SetTextDirection")
	purego.RegisterLibFunc(&ttfSetTextEngine, lib, "TTF_SetTextEngine")
	purego.RegisterLibFunc(&ttfSetTextFont, lib, "TTF_SetTextFont")
	purego.RegisterLibFunc(&ttfSetTextPosition, lib, "TTF_SetTextPosition")
	purego.RegisterLibFunc(&ttfSetTextScript, lib, "TTF_SetTextScript")
	purego.RegisterLibFunc(&ttfSetTextString, lib, "TTF_SetTextString")
	purego.RegisterLibFunc(&ttfSetTextWrapWhitespaceVisible, lib, "TTF_SetTextWrapWhitespaceVisible")
	purego.RegisterLibFunc(&ttfSetTextWrapWidth, lib, "TTF_SetTextWrapWidth")
	purego.RegisterLibFunc(&ttfStringToTag, lib, "TTF_StringToTag")
	purego.RegisterLibFunc(&ttfTagToString, lib, "TTF_TagToString")
	purego.RegisterLibFunc(&ttfTextWrapWhitespaceVisible, lib, "TTF_TextWrapWhitespaceVisible")
	purego.RegisterLibFunc(&ttfUpdateText, lib, "TTF_UpdateText")
	purego.RegisterLibFunc(&ttfVersion, lib, "TTF_Version")
	purego.RegisterLibFunc(&ttfWasInit, lib, "TTF_WasInit")
}
