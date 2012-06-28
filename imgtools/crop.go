package imgtools

import (
  "image"
  "image/color"
  "image/draw"
)

func Crop(img image.Image, width int, height int, padding int) (fimg draw.Image, err error) {
  // For now assume top left is bg color
  // future maybe avg outer rows of pixels
  bgcolor := img.At(img.Bounds().Min.X, img.Bounds().Min.Y)
  logoh, logow := height-2*padding, width-2*padding
  logorat := float32(logow) / float32(logoh)

  interior := findLogo(img, bgcolor)
  interior.Max = interior.Max.Add(image.Pt(1,1))

  center := func(rect image.Rectangle) image.Point {
    return image.Point{(rect.Max.X - rect.Min.X) / 2, (rect.Max.Y - rect.Min.Y) / 2}
  }
  fimg = image.NewRGBA(image.Rect(0, 0, width, height))

  rc := center(fimg.Bounds())
  origrat := float32(interior.Dx()) / float32(interior.Dy())

  if logorat > origrat {
    logow = int(origrat * float32(logoh))
  } else {
    logoh = int(float32(logow) / origrat)
  }
  logoimg := Resize(img, interior, logow, logoh)

  logorect := image.Rect(0, 0, logoimg.Bounds().Dx(), logoimg.Bounds().Dy())
  logorect = logorect.Add(rc.Sub(image.Pt(logorect.Dx()/2, logorect.Dy()/2)))

  draw.Draw(fimg, fimg.Bounds(), &image.Uniform{bgcolor}, image.ZP, draw.Src)
  draw.Draw(fimg, logorect, logoimg, image.ZP, draw.Src)
  return
}

func findLogo(img image.Image, bg color.Color) image.Rectangle {
  ol, il := img.Bounds().Min.X, img.Bounds().Max.X
  ir, or := img.Bounds().Min.X, img.Bounds().Max.X
  ot, it := img.Bounds().Min.Y, img.Bounds().Max.Y
  ib, ob := img.Bounds().Min.Y, img.Bounds().Max.Y

  for ot+1 < it {
    y := (ot + it) / 2
    hit := false
    for x := ol; x < or; x++ {
      if significantlyDifferent(bg, img.At(x, y)) {
        it = y
        if ir < x {
          ir = x
        }
        if il > x {
          il = x
        }
        hit = true
        break
      }
    }
    if !hit {
      ot = y
    }
  }

  for or-1 > ir {
    x := (or + ir) / 2
    hit := false
    for y := ot; y < ob; y++ {
      if significantlyDifferent(bg, img.At(x, y)) {
        ir = x
        if ib < y {
          ib = y
        }
        hit = true
        break
      }
    }
    if !hit {
      or = x
    }
  }
  for ob > ib+1 {
    y := (ob + ib) / 2
    hit := false
    for x := ol; x < or; x++ {
      if significantlyDifferent(bg, img.At(x, y)) {
        ib = y
        if il > x {
          il = x
        }
        hit = true
        break
      }
    }
    if !hit {
      ob = y
    }
  }
  for ol+1 < il {
    x := (ol + il) / 2
    hit := false
    for y := ot; y < ob; y++ {
      if significantlyDifferent(bg, img.At(x, y)) {
        il = x
        hit = true
        break
      }
    }
    if !hit {
      ol = x
    }
  }
  return image.Rect(il, it, ir, ib)
}

// Implementation doesn't get solid background colors perfectly
// uniform assuming for now that a two colors are significatly different
// if any one value is off by 5.
func significantlyDifferent(c1 color.Color, c2 color.Color) bool {
  stdrgba := func(c color.Color) (r, g, b, a uint32) {
    r, g, b, a = c.RGBA()
    r, g, b, a = r>>8, g>>8, b>>8, a>>8
    return
  }
  r1, g1, b1, a1 := stdrgba(c1)
  r2, g2, b2, a2 := stdrgba(c2)

  check := func(x, y uint32) bool {
    return x-10 < y && y < x+10
  }
  return !(check(r1, r2) && check(g1, g2) && check(b1, b2) && check(a1, a2))
}
