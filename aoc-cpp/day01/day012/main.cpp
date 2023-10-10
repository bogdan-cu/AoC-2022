#include <fstream>
#include <iostream>
#include <string>

int main()
{
    std::ifstream file;
    file.open("day1.txt");

    if (file.is_open())
    {
        std::string line;
        std::vector<int> loads;
        int total_weight = 0;
        int max_weight = 0;

        while (getline(file, line))
        {
            if (line == "")
            {
                loads.push_back(total_weight);
                total_weight = 0;
                continue;
            }

            try
            {
                total_weight += stoi(line);
            }
            catch (const std::exception &e)
            {
                std::cerr << e.what() << std::endl;
            }
            line.clear();
        }

        std::sort(loads.begin(), loads.end());
        for (auto i = loads.end() - 1; i >= loads.end() - 3 && i >= loads.begin(); i--)
        {
            max_weight += *i;
        }
        std::cout << "the strongest 3 lads carry a total of " << max_weight << std::endl;
    }
    else
    {
        std::cout << "could not open file" << std::endl;
    }

    return 0;
}
