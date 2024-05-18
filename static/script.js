function callApi() {
  fetch("/api/blogs")
    .then((response) => response.json())
    .then((data) => {
      console.log(data);
      document.getElementById("blogs").innerHTML = data.map((blog) => { return `<li>${blog.title}</li>` })
    });
}
