import { Box, useMediaQuery } from "@mui/material";
import Row1 from "./Row1";

const gridTemplateLargeScreens = `
  "a b c"
  "a b c"
  "a b c"
  "d e f"
  "d e f"
  "g h i"
  "g h i"
  "g h j"
  "g h j"
`;
const gridTemplateSmallScreens = `
  "a"
  "a"
  "a"
  "b"
  "b"
  "b"
  "c"
  "c"
  "c"
  "d"
  "d"
  "d"
  "e"
  "e"
  "f"
  "f"
  "f"
  "g"
  "g"
  "g"
  "h"
  "h"
  "h"
  "h"
  "i"
  "i"
  "j"
  "j"
`;

const Dashboard = () => {
  const isAboveMediumScreens = useMediaQuery("(min-width: 1200px)");
  return (
    <Box
      width="100%"
      height="100%"
      display="grid"
      gap="1.5rem"
      sx={
        isAboveMediumScreens
          ? {
              gridTemplateColumns: "repeat(3, minmax(250px, 2fr))", // Menyesuaikan lebar kolom untuk layar besar
              gridTemplateRows: "repeat(9, minmax(150px, 1fr))", // Menyesuaikan tinggi baris untuk layar besar
              gridTemplateAreas: gridTemplateLargeScreens,
            }
          : {
              gridAutoColumns: "1fr",
              gridAutoRows: "200px", // Menyesuaikan tinggi baris untuk layar kecil
              gridTemplateAreas: gridTemplateSmallScreens,
            }
      }
    >
      <Row1 />
    </Box>
  );
};

export default Dashboard;
