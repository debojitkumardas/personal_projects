#include <string>
#include <vector>
#include <iostream>

std::vector<std::string> Tokenise(std::string csv_line, char separator, char escape_seq = ' ') {
    std::vector<std::string> tokens;
    std::string token;
    auto start = csv_line.find_first_not_of(separator, 0);
    auto end = csv_line.find_first_of(separator, start);
    auto len = csv_line.length();

    if (escape_seq == ' ') {
        do {
            end = csv_line.find_first_of(separator, start);

            if (start == end) {
                if (start >= len) break;
                else {
                    start = end + 1;
                    continue;
                }
            }

            if (end != std::string::npos) token = csv_line.substr(start, end - start);
            else token = csv_line.substr(start, len - start);

            tokens.push_back(token);

            start = end + 1;
        } while (end != std::string::npos);
    }
    else {
        std::cout << "Code is in progress\n";
    }

    return tokens;
}
