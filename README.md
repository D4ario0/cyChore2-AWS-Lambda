# Cychore2
## About Cychore
Cychore, a short for *`cyclic chores`*, is a project develop by me and used by... me. Living with roomates? delegating tasks with equality? Sounds like familiar problem? Like every other engineer I grabbed a drawing board and quickly dished-out a spreadsheet with tasks for everyone, making sure everyone had to do all tasks from the list by the end of the month. 

Fast-forward 2 months into the system and it was about to fail, why? a bunch of college students easily forgetting to read their spreeadsheet shared in `Google Drive`. My inner engineer quickly tried to save the day creating the first iteration of cyChore, a `python` script that reads a `config` file, reads a list of users from a `json` and parses an `html` template to send everyone a plain email with their task for the week. Done. Simple. Effective.

## cyChore2, What's different?
`cyChore2` is just a more experienced me reinventing the wheel while keeping it simple. I know `Golang` now (ez to deploy btw). And everything is hosted on the cloud now.
- [x] Email Sender API built using `Go` standard library and `AWS SDK`.
- [x] Serverless API using `AWS Services`, mainly `Lambda Functions`, `API Gateway` and `CloudWatch` for debugging.
- [x] Best practices on project structure.
- [x] Email-focused HTML templates.
- [ ] Tests (On it)

Before I would open the project wait for it load in my IDE and click Run button vs. Now I double click a `.bat` file with a curl command.

## Discussion
If you are in the marketing business that sends a big amount of marketing and follow-up emails to your customer-base, it is worth considering the costs of building your own infrastructure as a microservice in contrast to use Sass products like `SendGrid` or `MailGun`.


 Thank you for the reading, **HAPPY CODING**
