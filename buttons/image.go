package buttons

import (
	"github.com/disintegration/gift"
	streamdeck "github.com/magicmonkey/go-streamdeck"
	"image"
)

// ImageFileButton represents a button with an image on it, where the image is loaded
// from a file.
type ImageButton struct {
	filePath      string
	img           image.Image
	updateHandler func(streamdeck.Button)
	btnIndex      int
	actionHandler streamdeck.ButtonActionHandler
}

// GetImageForButton is the interface implemention to get the button's image as an image.Image
func (btn *ImageButton) GetImageForButton(btnSize int) image.Image {
	// Resize the image to what the button wants
	g := gift.New(gift.Resize(btnSize, btnSize, gift.LanczosResampling))
	newimg := image.NewRGBA(image.Rect(0, 0, btnSize, btnSize))
	g.Draw(newimg, btn.img)
	return newimg
}

// SetButtonIndex is the interface implemention to set which button on the Streamdeck this is
func (btn *ImageButton) SetButtonIndex(btnIndex int) {
	btn.btnIndex = btnIndex
}

// GetButtonIndex is the interface implemention to get which button on the Streamdeck this is
func (btn *ImageButton) GetButtonIndex() int {
	return btn.btnIndex
}

// SetFilePath allows the image file to be changed on the fly
func (btn *ImageButton) SetImage(img image.Image) error {
	btn.img = img
	btn.updateHandler(btn)
	return nil
}

// RegisterUpdateHandler is the interface implemention to let the engine give this button a callback to
// use to request that the button image is updated on the Streamdeck.
func (btn *ImageButton) RegisterUpdateHandler(f func(streamdeck.Button)) {
	btn.updateHandler = f
}

// SetActionHandler allows a ButtonActionHandler implementation to be
// set on this button, so that something can happen when the button is pressed.
func (btn *ImageButton) SetActionHandler(a streamdeck.ButtonActionHandler) {
	btn.actionHandler = a
}

// Pressed is the interface implementation for letting the engine notify that the button has been
// pressed.  This hands-off to the specified ButtonActionHandler if it has been set.
func (btn *ImageButton) Pressed() {
	if btn.actionHandler != nil {
		btn.actionHandler.Pressed(btn)
	}
}

// NewImageFileButton creates a new ImageFileButton with the specified image on it
func NewImageButton(image image.Image) *ImageButton {
	return &ImageButton{img: image}
}
