# Todo

A simple tool designed to facilitate the synchronization of text content between local and remote servers.

### Features

- **Push & Pull:** Todo allows users to push local text content to update the server-side data and pull remote data to update local content efficiently.
- **Configurability:** Users can configure settings through a `config.yaml` file to tailor Todo to their specific needs.
- **Deployment:** The project can be compiled and deployed easily using the provided Makefile. Building for Linux AMD64 architecture is achieved with `make linux-amd64`, while deployment can be initiated using `make deploy`.

## Usage

1. **Configuration:**
   Edit the `config.yaml` file according to `config/config.go`.

2. **Pushing Content:**
   To update server-side data with local changes, use the following command:
   ```
   cat todo.md | todo-client push
   ```

3. **Pulling Content:**
   To update local content with changes from the server, execute:
   ```
   ./todo-client pull > todo.md
   ```

## License

Todo is licensed under the [GNU General Public License Version 3 (GPL V3)](https://www.gnu.org/licenses/gpl-3.0.en.html). Please see the `LICENSE` file for more details.

## Contributions

Contributions to Todo are welcome! Feel free to submit bug reports, feature requests, or pull requests to help improve the project. Let's build something great together!