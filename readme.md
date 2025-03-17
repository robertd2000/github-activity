# GitHub User Activity

Solution for the [GitHub User Activity](https://github.com/robertd2000/github-activity) challenge from [roadmap.sh](https://roadmap.sh/) (without any external libraries).

## How to run

Clone the repository and run the following command:

```bash
git clone https://github.com/robertd2000/github-activity.git
cd github-activity
```

Run the following command to run the project:

```bash

# to get all user activities
go run .  <username>

#example
go run . robertd2000

# to get only one type of activities:
go run . robertd2000 DeleteEvent # return only DeleteEvent activities

```

[https://roadmap.sh/projects/github-user-activity](https://roadmap.sh/projects/github-user-activity)
