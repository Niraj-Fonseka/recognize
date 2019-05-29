# RECOGNIZE


Webserver that streams data when a human face is detected using gRPC streaming.

Clients can subscribe to the server and read the data. Currently it's outputting the cartesian corrdinates of where the face was detected in the frame.

The face recognition is done using OpenCV's haarcascades. I haven't really had the time to test for the false positve rate. But it's in progress. I'm also currently in the process of converting and improving the facial detection to actuation facial identification using OpenCV's built in dnn ( Deep neural network ) module. When this is implemented you will be able to identify certian people after training a model with a few pictures of that person. 



## Installation and setup 


### Running the server
```
(1) Install the virtualenv
        pip install virtualenv

(2) create a virtualenv and activate it
        source /virtualenvname/bin/activate

(3) install dependencies 
        pip install -r requirements.txt

(4) run the server 
        python server.py
```

Output should be similar to 
```
python server.py
Detector watching ..
```

### Running the example go client 
```
(1) Make sure you have go installed ( > go 1.11)

(2) Enable go modules 
    export GO111MODULE=on

(3) Run the client
    go run client.go
```

The client will call the Detect function in the server. Get in front of the camera ! If you have a light near your camera to show its activated,
it will turn on. 

In the client you will see an output similar to this 

```
2019/05/29 00:44:40 X : 278 , Y : 44 , W : 228 , H : 228 
2019/05/29 00:44:42 X : 292 , Y : 50 , W : 228 , H : 228 
2019/05/29 00:44:45 X : 309 , Y : 55 , W : 228 , H : 228 
2019/05/29 00:44:45 X : 304 , Y : 55 , W : 228 , H : 228 
2019/05/29 00:44:52 X : 314 , Y : 70 , W : 228 , H : 228 
2019/05/29 00:44:52 X : 355 , Y : 86 , W : 228 , H : 228 
2019/05/29 00:44:54 X : 384 , Y : 50 , W : 228 , H : 228 
2019/05/29 00:44:54 X : 389 , Y : 59 , W : 228 , H : 228 
```