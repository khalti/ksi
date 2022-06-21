# Khalti Sprint Interface (KSI)

KSI (Khalti Sprint Interface) is a sprint-flow-interface (SFI {made up}) which can be used to control the sprint flow and know the sprint status (repo. wise) via CLI.

KSI can automate tasks such as:
- **Starting/Ending Sprint:** With KSI’s sprint command options, you can start/end current sprint or old sprints which will cleanup the branches (remotely and locally) and merge code (locally).
- **Basic git jazz:** Every other main git commands are restyled with KSI.
- **What?:** With “What” command options, developer can see what sprint is running, the branches related to the sprint, current branch and what not.

## Commands and stories
> Work in progress and subject to change

#### sprint (`ksi sprint`)

This command will basically fetch all of the branches and organize branches w.r.t. sprint branches to get started for the first time.

##### fetch (`ksi sprint fetch`)
This command will basically fetch all of the branches and organize branches w.r.t. sprint branches to get started for the first time.

##### start (`ksi sprint start`)
This sub-command will create a new sprint branch, pull from the old sprint branch (or you can create new sprint branch from any branch as source of truth), and push the codes to get started for new sprint.

The new sprint branch should look like `sp06180624`

##### end (`ksi sprint end`)
This command will end the current sprint. While doing it, it will also delete all of the possible sprint related feature branches remotely and locally. 

####  feature (`ksi feature`)
This command will deals with all of the jazz related to feature branches.

##### new (`ksi feature new`)
This sub-command will create a new feature branch associating with the current sprint. You can specify feature name via name input prompt. This command will prompt you with options to choose from to specify the nature of branch.

Natures of a branch could be:

- Feature
- Enhancement
- Bug Fix
- Chore
- Other

##### pull (`ksi feature pull`)
This sub-command will pull all of the latest code (remote) to the current feature branch (local). You can, otherwise, also specify the some other branch as source of truth.

##### stage (`ksi feature stage`)
This sub-command is basically git add . and git commit ... but restyled. It will stage the changes locally and while doing it, it will ask developer to add commit title and description.

##### push (`ksi feature push`)
This sub-command will initiate the push protocol, where developers will be asked few questions which will prepare a push and execute it.

This  sub-command will try to pull any latest changes (remote) and only if there are no any new changes or doesn’t result in conflicts, it will continue or else it will exit for developer to resolve the conflicts.

Then it will run git add . command to stage all of the changes.
> If there are new changes, it will prompt developer to enter the commit title. Commit title can be 50 characters long. So, be wise while writing title. You can add other info about commit into description which will prompt after title is entered.

Then finally, will push the commits.
