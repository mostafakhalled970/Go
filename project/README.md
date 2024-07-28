# Go Simple Main

<!-- ![alt text](https://miro.medium.com/v2/resize:fit:1400/format:webp/1*HtCjHzGwf6iWNqXu5Cndsg.png) -->
![alt text](https://res.cloudinary.com/practicaldev/image/fetch/s--KFzYj8KY--/c_imagga_scale,f_auto,fl_progressive,h_420,q_auto,w_1000/https://thepracticaldev.s3.amazonaws.com/i/jz1v8e2wrog01wo0ljyn.jpg)
this project is focusing on gitlab CI, You could also use it as <br > a template to initiate your projects

## How it works
there is a file named `.gitlab-ci.yml` which deals with gitlab CI.
whenever you push your code to your gitlab repository it checks two stages: 
* Build
* Test:
  * Test stage relates to calc package in the project which has a `SUM()` function and has a `calc_test.go` file that test multiple test cases. **If test cases goes successful means it passed the test stage**.


## Requirements

* docker
    * docker-compose

## Dependencies

* Gin-Gonic

## Usage

```bash
docker build --tag go_simple_main:1.0.0 .
docker run -p 8080:8080 go_simple_main:1.0.0
OR
docker-compose up -d
```

## You could also test it by doing:

```bash
curl -X GET http://127.0.0.1:8080/ping
```

## Contributing

Pull requests are welcome. For major changes, please open an issue first
to discuss what you would like to change.

Please make sure to update tests as appropriate.
