## Quick Start Instructions for CWE REST API Users

The API is available without any need to register or use any credentials.

The URL to access the API is: https://cwe-api.mitre.org/api/v1/cwe/

The available endpoints are:

To view detailed documentation of the endpoints, use the swagger.io website:

Swagger editor and import the file openapi.json

Note that a request for the children, parents, descendants or ancestors of an existing CWE will return an empty list, [], and a status of 200 if none exist.  However, requesting a CWE id that does not exist from one of those endpoints will return a status of 404.

Additionally, using an endpoint for the wrong type of CWE (e.g., using the weakness endpoint for a category) will also return a status of 404.
 
Here are examples of each endpoint to try:

Note: examples use the full URL
 
- https://cwe-api.mitre.org/api/v1/cwe/version
- https://cwe-api.mitre.org/api/v1/cwe/74,79
- https://cwe-api.mitre.org/api/v1/cwe/weakness/79
- https://cwe-api.mitre.org/api/v1/cwe/category/189
- https://cwe-api.mitre.org/api/v1/cwe/view/1425
- https://cwe-api.mitre.org/api/v1/cwe/74/parents?view=1000
- https://cwe-api.mitre.org/api/v1/cwe/74/descendants?view=1000
- https://cwe-api.mitre.org/api/v1/cwe/74/children?view=1000
- https://cwe-api.mitre.org/api/v1/cwe/74/ancestors?view=1000

