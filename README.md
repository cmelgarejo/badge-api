# Badge API

## This is an implementation of the following

> Key things we look for
>
> Breaking down problems and solution design.
> Documenting system design decisions.
> Code quality and testability.
> Demonstration of the breadth of knowledge, and depth in one or two areas in the given time.
>
> The project
> Cuely is about to build its gamification system that supports different types of rewards, e.g. points, badges, levels, titles, or awards.
>
> Let’s build a backend REST API to support creating, updating, and retrieving rewards for the authenticated user.
>
> It’s up to you to decide on how to authenticate the user or verify the user’s identity.
>
> `/v1/point/user_id`
>
> GET - params (start_date, end_date) - Get points users received in a time period. If missing, return historic points
> POST - params (points, date, metadata) - Add points to the authenticated user
>
> `/v1/badge/org_id`
>
> GET - Get all badges available to an organization
> POST - params (name, image) Create a badge
>
> `/v1/badge/user_id`
>
> GET - Get badges users received with (id, name, and image)
> POST - params (badge id) - Add a badge to the user
>
> If you want to build an interactive UI to demonstrate the usage of the API, that’s a plus but not required.
>
> The review meeting
> Be ready to discuss the design trade-offs and how to tweak the servers to add more capabilities.

## Some changes I've made to the implementation

### Resources

Namely:

- `/v1/badge/org_id` -> `/v1/badges/org/org_id`
- `/v1/badge/user_id` -> `/v1/badges/user/user_id`

Because discriminating the id's either user or org within the same endpoint its a bit hard to do, and I think it's better to have a separate endpoint for each.

## Running the project

Gotta have a postgresql instance running (either local or dockerized) heck I'll put a docker one right now. Done.

run `docker compose up -d` on the directory and you're set for persistance

Then you have to run the following:

`make`

that's it. :)

You can now use the POSTMAN collection to test the API.

## TODO

- Tests - build mockdb on the `internal/data` folder and run the tests there.
- Frontend - build a frontend on the `frontend` folder. A NextJS app would be nice.
