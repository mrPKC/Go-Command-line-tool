Family Tree CLI
Family Tree CLI is a Go-based command-line application for managing family tree data.

Installation
Make sure you have Go installed. Clone the repository and run:



go install

Build Command : go build -o family-tree
cp family-tree /usr/local/bin
Usage
Add Person


family-tree add person <name> <gender> <dob>
Add a new person to the family tree with the specified name, gender, and date of birth.

Get Person List


family-tree get person list
Retrieve a list of all people in the family tree.

Get Person Details


family-tree get person <name>
Retrieve details of a specific person in the family tree by providing their name.

Add Relationship


family-tree add relationship <name>
Add a new relationship to the family tree for the specified person.

Get Relationship List


family-tree get relationship list
Retrieve a list of all relationships in the family tree.

Get Relationship Details


family-tree get relationship <name>
Retrieve details of a specific relationship in the family tree by providing the person's name.

Connect People


family-tree connect <name> as <relation> of <name>
Connect two people in the family tree with the specified relationship.

Count Relationships


family-tree count <relation> of <name>
Count the number of relationships of a specific type for a given person.

List Relationships


family-tree <relation> of <name>
List all people related to a specific person by the specified relationship.

License
This project is licensed under the MIT License - see the LICENSE file for details.

Replace <name>, <gender>, <dob>, <relation>, and other placeholders with actual values or parameters. Update the installation instructions based on your project structure and dependencies. Additionally, include information about licensing if applicable.

Feel free to expand and customize this template according to your specific project details.