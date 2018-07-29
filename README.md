this is the first part, the api.

It's clearly not the best possible solution, as it's lacking several features i would say mandatory in a modern api, like caching, logging, authentication, and a proper error handling (the api always returns 200, even in case of errors).

However, given that I had to learn from zero the basics of both the frontend and the backend technologies used, in a weekend, I can say I am quite proud of the result.

The api it's done using a package called mux as a router and it consist mostly of simple map manipulation,using a service that goes trough 2 in-memory repositories, one for the questions and one for the use user results.

There are 3 endpoints : a GET /questions to retrieve the list of questions toghether with the options, a POST /answers to send back to the server the answers and the userId, and a GET users/<ID>/results to return some basic info about the performance of the user, given that the user exists.

I also added a very simple CORS manager in order to avoid issues while using the frontend running locally.

running the .exe file created after the build will put the listener at the port 8000 of localhost, and it is queryable both using the frontend in the quizUI repository or with Postman.

let me know for any question the code may arise and I will try to answer as soon as possible.

thank you!
