### Vue.js SPA + Golang API backend

This repository provides a simple demo web app using Vue.js and querying a backend written in Golang over XHR.

#### Usage
```bash
git clone https://github.com/Simon-L/vue-go.git
cd vue-go/vue
npm install
go run ../go/back.go & npm run dev
# stop the backend
killall back
```

#### Project structure

The `vue/` folder is mostly the [vue-webpack template](https://github.com/vuejs-templates/webpack) with the following added/modified:
* Bulma CSS framework
* `vue-resource` HTTP client using XMLHttpRequest
* SASS/SCSS pre-processor enabled
* `vue/config/index.js` includes configuration to proxy `/api` to the Golang backend [1]

Even the component name `Hello.vue` is kept from the original boilerplate :)

The `go/` folder contains `back.go`, a minimal program with no dependencies serving the current time in JSON to a request on route `/`.

**Notes**  
[1] The webapp is listening on `localhost:8080`, the go program listens on `localhost:9000`, the proxy configuration forwards everything from `localhost:8080/api` to `localhost:9000/`, hence why the HTTP request to `localhost:8080/api` is received in the golang backend on `/` and receives the current time in response.
