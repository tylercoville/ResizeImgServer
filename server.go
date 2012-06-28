package main

import (
  "image"
  "image/jpeg"
  "image/png"
  "net/http"
  "resizeimgserver/imgtools"
  "strconv"
)

func main() {
  http.HandleFunc("/logoresize/", handler)
  http.ListenAndServe(":8080", nil)
}

func handler(w http.ResponseWriter, r *http.Request) {
  imgUrl := r.FormValue("image")
  width, errw := strconv.Atoi(r.FormValue("width"))
  height, errh := strconv.Atoi(r.FormValue("height"))
  padding, errp := strconv.Atoi(r.FormValue("padding"))
  resp, err := http.Get(imgUrl)

  if errw != nil || errh != nil || errp != nil || err != nil {
    http.Error(w, "Invalid parameters", 400)
    return
  }

  img, fmt, err := image.Decode(resp.Body)
  if err != nil { http.Error(w, err.Error(), 400); return }

  defer resp.Body.Close()
  fimg, _ := imgtools.Crop(img, width, height, padding)

  switch fmt {
    case "png":
      png.Encode(w, fimg)
    case "jpeg":
      jpeg.Encode(w, fimg, nil)
  }

}
