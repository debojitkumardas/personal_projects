// NOLINTBEGIN
// NOLINTEND
#include "helper.hpp"

int main() {

    char ch = 'y';
    Helper app{};

    if (ch == 'y')
        app.Init("data/20200317.csv");

    return 0;
}
