FROM jamesandersen/alpine-golang-opencv3:edge
RUN apk --no-cache add git

WORKDIR /go/src/github.com/ryanbmilbourne/syr-sudoku-backend
COPY . .

RUN go get github.com/gin-gonic/gin
RUN go get github.com/sirupsen/logrus
RUN go get github.com/Wrenky/sudoKu

#RUN go-wrapper download   # "go get -d -v ./..."
#RUN go-wrapper install    # "go install -v ./..."

RUN go build -o srv ./cmd/srv

FROM alpine:edge
RUN echo "http://dl-cdn.alpinelinux.org/alpine/edge/testing" >> /etc/apk/repositories
RUN apk --no-cache add ca-certificates opencv-libs
RUN ln /usr/lib/libopencv_core.so.3.2.0 /usr/lib/libopencv_core.so \
    && ln /usr/lib/libopencv_highgui.so.3.2.0 /usr/lib/libopencv_highgui.so \
    && ln /usr/lib/libopencv_imgcodecs.so.3.2.0 /usr/lib/libopencv_imgcodecs.so \
    && ln /usr/lib/libopencv_imgproc.so.3.2.0 /usr/lib/libopencv_imgproc.so \
    && ln /usr/lib/libopencv_ml.so.3.2.0 /usr/lib/libopencv_ml.so \
    && ln /usr/lib/libopencv_objdetect.so.3.2.0 /usr/lib/libopencv_objdetect.so \
    && ln /usr/lib/libopencv_photo.so.3.2.0 /usr/lib/libopencv_photo.so
WORKDIR /root/

# ** MAGIC HAPPENS HERE! **
# We can grab the binary built in the previous stage and copy it to our "clean" image
COPY --from=0 /go/src/github.com/ryanbmilbourne/syr-sudoku-backend/srv .
#COPY web ./web
# Set the PORT environment variable inside the container
ENV PORT 8080
# Expose port 5000 to the host so we can access our application
EXPOSE 8080
CMD ["./srv"]  
