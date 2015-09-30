package keyboard

import "image"

func (k Keyboard) SetImage(img image.Gray) {
	packet := make([]byte, 992)
	packet[0] = 0x03
	imgOffset := 32

	var curr, row byte
	for offset := 0; offset < 5; offset++ {
		for col := 0; col < 160; col++ {
			curr = 0
			for row = 0; row < 8; row++ {
				pixOffset := (offset*8+int(row)-img.Rect.Min.Y)*img.Stride + (col - img.Rect.Min.X)
				if img.Pix[pixOffset]>>4 > 0 {
					curr += 1 << row
				}
			}
			packet[imgOffset] = curr
			imgOffset++
		}
	}
	k.Device.Write(packet)
}
