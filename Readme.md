
<!-- PROJECT LOGO -->
<br />
<div align="center">
<h3 align="center">Smallest amount of items or Smallest package finder</h3>

  <p align="center">
    An awesome project to find the smallest amount of items or the smallest package to fit a given amount of items.
    <br />
    <br />
    <a href="https://github.com/github_username/repo_name">View Demo</a>
    &middot;
    <a href="https://github.com/github_username/repo_name/issues/new?labels=bug&template=bug-report---.md">Report Bug</a>
    &middot;
    <a href="https://github.com/github_username/repo_name/issues/new?labels=enhancement&template=feature-request---.md">Request Feature</a>
  </p>
</div>



<!-- TABLE OF CONTENTS -->
<details>
  <summary>Table of Contents</summary>
  <ol>
    <li>
      <a href="#about-the-project">About The Project</a>
      <ul>
        <li><a href="#built-with">Built With</a></li>
      </ul>
    </li>
    <li>
      <a href="#getting-started">Getting Started</a>
      <ul>
        <li><a href="#prerequisites">Prerequisites</a></li>
        <li><a href="#installation">Installation</a></li>
        <li><a href="#unit-tests">Unit Tests</a></li>
        <li><a href="#acceptance-tests">Acceptance Tests</a></li>
      </ul>
    </li>

  </ol>
</details>



<!-- ABOUT THE PROJECT -->
## About The Project
This project is a simple example of how to find the smallest amount of items or the smallest package to fit a given amount of items.
It follows the following rules:

1. Only whole packs can be sent. Packs cannot be broken open.
2. Within the constraints of Rule 1 above, send out the least amount of items to fulfil the order.
3. Within the constraints of Rules 1 & 2 above, send out as few packs as possible to fulfil each order.
   (Please note, rule #2 takes precedence over rule #3)


The project is divided into two parts:
1. A backend written in Golang that exposes a REST API to calculate the smallest amount of items or the smallest package to fit a given amount of items.
2. A frontend written in Next.js that consumes the REST API and provides a simple UI to interact with the backend.


<p align="right">(<a href="#readme-top">back to top</a>)</p>



### Built With

* [![Next][Next.js]][Next-url]
* [![React][React.js]][React-url]
* Golang

<p align="right">(<a href="#readme-top">back to top</a>)</p>



<!-- GETTING STARTED -->
## Getting Started

### Installation

1. Clone the repo
   ```sh
   git clone https://github.com/vcayetano/fulfillment.git
   ```
2. Run the project
   ```sh
   make run
   ```

<p align="right">(<a href="#readme-top">back to top</a>)</p>

### Unit Tests

1. Run the unit tests
   ```sh
   make unit-tests
   ```
2. Results will be displayed in the terminal
```code
go test -v ./specifications/...
=== RUN   TestPack
=== RUN   TestPack/It_returns_the_correct_number_of_items_for_the_edge_cases
=== RUN   TestPack/It_returns_the_correct_number_of_items_for_the_edge_cases/returns_correct_pack_count_when_order_count_is_500000_(edge_case)
=== RUN   TestPack/it_returns_the_correct_number_of_items_to_use_for_the_order_count
=== RUN   TestPack/it_returns_the_correct_number_of_items_to_use_for_the_order_count/returns_correct_pack_count_when_order_count_is_0
=== RUN   TestPack/it_returns_the_correct_number_of_items_to_use_for_the_order_count/returns_correct_pack_count_when_order_count_is_1
=== RUN   TestPack/it_returns_the_correct_number_of_items_to_use_for_the_order_count/returns_correct_pack_count_when_order_count_is_250
=== RUN   TestPack/it_returns_the_correct_number_of_items_to_use_for_the_order_count/returns_correct_pack_count_when_order_count_is_251
=== RUN   TestPack/it_returns_the_correct_number_of_items_to_use_for_the_order_count/returns_correct_pack_count_when_order_count_is_500
=== RUN   TestPack/it_returns_the_correct_number_of_items_to_use_for_the_order_count/returns_correct_pack_count_when_order_count_is_501
=== RUN   TestPack/it_returns_the_correct_number_of_items_to_use_for_the_order_count/returns_correct_pack_count_when_order_count_is_700
=== RUN   TestPack/it_returns_the_correct_number_of_items_to_use_for_the_order_count/returns_correct_pack_count_when_order_count_is_1000
=== RUN   TestPack/it_returns_the_correct_number_of_items_to_use_for_the_order_count/returns_correct_pack_count_when_order_count_is_1001
=== RUN   TestPack/it_returns_the_correct_number_of_items_to_use_for_the_order_count/returns_correct_pack_count_when_order_count_is_1500
=== RUN   TestPack/it_returns_the_correct_number_of_items_to_use_for_the_order_count/returns_correct_pack_count_when_order_count_is_2000
=== RUN   TestPack/it_returns_the_correct_number_of_items_to_use_for_the_order_count/returns_correct_pack_count_when_order_count_is_4000
=== RUN   TestPack/it_returns_the_correct_number_of_items_to_use_for_the_order_count/returns_correct_pack_count_when_order_count_is_12500
=== RUN   TestPack/it_returns_the_correct_number_of_items_to_use_for_the_order_count/returns_correct_pack_count_when_order_count_is_12001
--- PASS: TestPack (0.00s)
--- PASS: TestPack/It_returns_the_correct_number_of_items_for_the_edge_cases (0.00s)
--- PASS: TestPack/It_returns_the_correct_number_of_items_for_the_edge_cases/returns_correct_pack_count_when_order_count_is_500000_(edge_case) (0.00s)
--- PASS: TestPack/it_returns_the_correct_number_of_items_to_use_for_the_order_count (0.00s)
--- PASS: TestPack/it_returns_the_correct_number_of_items_to_use_for_the_order_count/returns_correct_pack_count_when_order_count_is_0 (0.00s)
--- PASS: TestPack/it_returns_the_correct_number_of_items_to_use_for_the_order_count/returns_correct_pack_count_when_order_count_is_1 (0.00s)
--- PASS: TestPack/it_returns_the_correct_number_of_items_to_use_for_the_order_count/returns_correct_pack_count_when_order_count_is_250 (0.00s)
--- PASS: TestPack/it_returns_the_correct_number_of_items_to_use_for_the_order_count/returns_correct_pack_count_when_order_count_is_251 (0.00s)
--- PASS: TestPack/it_returns_the_correct_number_of_items_to_use_for_the_order_count/returns_correct_pack_count_when_order_count_is_500 (0.00s)
--- PASS: TestPack/it_returns_the_correct_number_of_items_to_use_for_the_order_count/returns_correct_pack_count_when_order_count_is_501 (0.00s)
--- PASS: TestPack/it_returns_the_correct_number_of_items_to_use_for_the_order_count/returns_correct_pack_count_when_order_count_is_700 (0.00s)
--- PASS: TestPack/it_returns_the_correct_number_of_items_to_use_for_the_order_count/returns_correct_pack_count_when_order_count_is_1000 (0.00s)
--- PASS: TestPack/it_returns_the_correct_number_of_items_to_use_for_the_order_count/returns_correct_pack_count_when_order_count_is_1001 (0.00s)
--- PASS: TestPack/it_returns_the_correct_number_of_items_to_use_for_the_order_count/returns_correct_pack_count_when_order_count_is_1500 (0.00s)
--- PASS: TestPack/it_returns_the_correct_number_of_items_to_use_for_the_order_count/returns_correct_pack_count_when_order_count_is_2000 (0.00s)
--- PASS: TestPack/it_returns_the_correct_number_of_items_to_use_for_the_order_count/returns_correct_pack_count_when_order_count_is_4000 (0.00s)
--- PASS: TestPack/it_returns_the_correct_number_of_items_to_use_for_the_order_count/returns_correct_pack_count_when_order_count_is_12500 (0.00s)
--- PASS: TestPack/it_returns_the_correct_number_of_items_to_use_for_the_order_count/returns_correct_pack_count_when_order_count_is_12001 (0.00s)
PASS
ok      github.com/vcayetano/fulfillment/specifications (cached)

```


<p align="right">(<a href="#readme-top">back to top</a>)</p>

### Acceptance Tests
Testcontainers is used to run acceptance tests. The acceptance tests are written in Golang and uses the testcontainers library to run the tests in a containerized environment. The acceptance tests will test the REST API exposed by the backend.

1. Run the unit tests
   ```sh
   make acceptance-test-http
   ```
2. Results will be displayed in the terminal
```code
go test -v ./cmd/httpserver/...
?       github.com/vcayetano/fulfillment/cmd/httpserver [no test files]
=== RUN   TestPackerServer
2025/01/19 18:58:18 github.com/testcontainers/testcontainers-go - Connected to docker: 
  Server Version: 24.0.6
  API Version: 1.43
  Operating System: Docker Desktop
  Total Memory: 8938 MB
  Testcontainers for Go Version: v0.35.0
  Resolved Docker Host: unix:///Users/richard/.docker/run/docker.sock
  Resolved Docker Socket Path: /var/run/docker.sock
  Test SessionID: 979cc6ce4a948015618db0ab4b2fbb18effa9b6bff8edea3b6250be5baf17828
  Test ProcessID: d3c0378b-5cc0-4353-ac78-95a13bcf7931
2025/01/19 18:58:18 🐳 Building image 6fd51048-3992-48a0-bad7-e1a15d97d040:83dd8de3-adf6-4dd4-aa41-c4b4f4496e0a
Step 1/13 : ARG GO_VERSION=1.23-alpine
Step 2/13 : FROM golang:$GO_VERSION
 ---> 466670fb5708
Step 3/13 : WORKDIR /app
 ---> Using cache
 ---> 0ea7f31818c4
Step 4/13 : COPY go.mod ./
 ---> ed9882261708
Step 5/13 : RUN go mod download
 ---> Running in 6950c784f9e0
Removing intermediate container 6950c784f9e0
 ---> e2f620063946
Step 6/13 : COPY . .
 ---> d1de098be005
Step 7/13 : RUN go build -o svr cmd/httpserver/*.go
 ---> Running in d18530c9ee2f
Removing intermediate container d18530c9ee2f
 ---> 256f1426c560
Step 8/13 : CMD [ "./svr" ]
 ---> Running in 25ab3eb809ca
Removing intermediate container 25ab3eb809ca
 ---> 1a085745af33
Step 9/13 : LABEL org.testcontainers=true
 ---> Running in ea2d447c7a8f
Removing intermediate container ea2d447c7a8f
 ---> 886e8e4ea778
Step 10/13 : LABEL org.testcontainers.lang=go
 ---> Running in 5de5d2438124
Removing intermediate container 5de5d2438124
 ---> 0909ec4a325f
Step 11/13 : LABEL org.testcontainers.reap=true
 ---> Running in 0d5d3219c8b8
Removing intermediate container 0d5d3219c8b8
 ---> 43e494306952
Step 12/13 : LABEL org.testcontainers.sessionId=979cc6ce4a948015618db0ab4b2fbb18effa9b6bff8edea3b6250be5baf17828
 ---> Running in 2d0677df1777
Removing intermediate container 2d0677df1777
 ---> 38fe7e1e62ac
Step 13/13 : LABEL org.testcontainers.version=0.35.0
 ---> Running in 0a02e475979e
Removing intermediate container 0a02e475979e
 ---> ac1c5700f491
Successfully built ac1c5700f491
Successfully tagged cc919f00-4f2b-4780-9b73-5c3332601ea9:277f2147-fbb7-452c-b077-a2f3cedb54ea
2025/01/19 18:58:38 ✅ Built image cc919f00-4f2b-4780-9b73-5c3332601ea9:277f2147-fbb7-452c-b077-a2f3cedb54ea
2025/01/19 18:58:38 🐳 Creating container for image cc919f00-4f2b-4780-9b73-5c3332601ea9:277f2147-fbb7-452c-b077-a2f3cedb54ea
2025/01/19 18:58:38 🐳 Creating container for image testcontainers/ryuk:0.11.0
2025/01/19 18:58:38 ✅ Container created: 6e39c7c9cc52
2025/01/19 18:58:38 🐳 Starting container: 6e39c7c9cc52
2025/01/19 18:58:38 ✅ Container started: 6e39c7c9cc52
2025/01/19 18:58:38 ⏳ Waiting for container id 6e39c7c9cc52 image: testcontainers/ryuk:0.11.0. Waiting for: &{Port:8080/tcp timeout:<nil> PollInterval:100ms skipInternalCheck:false}
2025/01/19 18:58:38 🔔 Container is ready: 6e39c7c9cc52
2025/01/19 18:58:38 ✅ Container created: e4947835a53c
2025/01/19 18:58:38 🐳 Starting container: e4947835a53c
2025/01/19 18:58:39 ✅ Container started: e4947835a53c
2025/01/19 18:58:39 ⏳ Waiting for container id e4947835a53c image: cc919f00-4f2b-4780-9b73-5c3332601ea9:277f2147-fbb7-452c-b077-a2f3cedb54ea. Waiting for: &{Port:8080 timeout:0x140002f9940 PollInterval:100ms skipInternalCheck:false}
2025/01/19 18:58:39 🔔 Container is ready: e4947835a53c
packages:  [{"size":23,"quantity":2},{"size":31,"quantity":7},{"size":53,"quantity":9429}]

2025/01/19 18:58:39 🐳 Stopping container: e4947835a53c
2025/01/19 18:58:39 ✅ Container stopped: e4947835a53c
2025/01/19 18:58:39 🐳 Terminating container: e4947835a53c
2025/01/19 18:58:39 🚫 Container terminated: e4947835a53c
--- PASS: TestPackerServer (21.60s)
PASS
ok      github.com/vcayetano/fulfillment/cmd/httpserver/tests   21.871s

```

<p align="right">(<a href="#readme-top">back to top</a>)</p>


<!-- MARKDOWN LINKS & IMAGES -->
<!-- https://www.markdownguide.org/basic-syntax/#reference-style-links -->
[contributors-shield]: https://img.shields.io/github/contributors/github_username/repo_name.svg?style=for-the-badge
[contributors-url]: https://github.com/github_username/repo_name/graphs/contributors
[forks-shield]: https://img.shields.io/github/forks/github_username/repo_name.svg?style=for-the-badge
[forks-url]: https://github.com/github_username/repo_name/network/members
[stars-shield]: https://img.shields.io/github/stars/github_username/repo_name.svg?style=for-the-badge
[stars-url]: https://github.com/github_username/repo_name/stargazers
[issues-shield]: https://img.shields.io/github/issues/github_username/repo_name.svg?style=for-the-badge
[issues-url]: https://github.com/github_username/repo_name/issues
[license-shield]: https://img.shields.io/github/license/github_username/repo_name.svg?style=for-the-badge
[license-url]: https://github.com/github_username/repo_name/blob/master/LICENSE.txt
[linkedin-shield]: https://img.shields.io/badge/-LinkedIn-black.svg?style=for-the-badge&logo=linkedin&colorB=555
[linkedin-url]: https://linkedin.com/in/linkedin_username
[product-screenshot]: images/screenshot.png
[Next.js]: https://img.shields.io/badge/next.js-000000?style=for-the-badge&logo=nextdotjs&logoColor=white
[Next-url]: https://nextjs.org/
[React.js]: https://img.shields.io/badge/React-20232A?style=for-the-badge&logo=react&logoColor=61DAFB
[React-url]: https://reactjs.org/
[Vue.js]: https://img.shields.io/badge/Vue.js-35495E?style=for-the-badge&logo=vuedotjs&logoColor=4FC08D
[Vue-url]: https://vuejs.org/
[Angular.io]: https://img.shields.io/badge/Angular-DD0031?style=for-the-badge&logo=angular&logoColor=white
[Angular-url]: https://angular.io/
[Svelte.dev]: https://img.shields.io/badge/Svelte-4A4A55?style=for-the-badge&logo=svelte&logoColor=FF3E00
[Svelte-url]: https://svelte.dev/
[Laravel.com]: https://img.shields.io/badge/Laravel-FF2D20?style=for-the-badge&logo=laravel&logoColor=white
[Laravel-url]: https://laravel.com
[Bootstrap.com]: https://img.shields.io/badge/Bootstrap-563D7C?style=for-the-badge&logo=bootstrap&logoColor=white
[Bootstrap-url]: https://getbootstrap.com
[JQuery.com]: https://img.shields.io/badge/jQuery-0769AD?style=for-the-badge&logo=jquery&logoColor=white
[JQuery-url]: https://jquery.com # fulfillment
