
# gamma_rate_list it's a list of dicts that will looks like this
# [
#   {
#       0: occurences of 0 in the first bit,
#       1: occurences of 1 in the first bit,
#   },
#   {
#       0: occurences of 0 in the second bit,
#       1: occurences of 1 in the second bit,
#   },
#   ...
#   ...
# ]
    
gamma_rate_list = []
gamma_rate = ""
epsilon_rate = ""

with open("../input.txt") as file:
    for line in file:
        for i, num in enumerate(line.rstrip()):
            try:
                if int(num) == 0:
                    gamma_rate_list[i]["0"] += 1
                else:
                    gamma_rate_list[i]["1"] += 1
            # this will fail (only) at the first iteration,
            # kinda looks more efficient than checking it every time
            except:
                if int(num) == 0:
                    gamma_rate_list.append({"0": 1, "1": 0})
                else:
                    gamma_rate_list.append({"0": 0, "1": 1})

for bit in gamma_rate_list:
    if bit["0"] > bit["1"]:
        gamma_rate += "0"
        epsilon_rate += "1"
    else:
        gamma_rate += "1"
        epsilon_rate += "0"

print(int(gamma_rate, 2) * int(epsilon_rate, 2))
