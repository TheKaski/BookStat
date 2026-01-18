import ShelvesPreviewContainer from "../components/ShelvesPreviewContainer";
import UnassignedBooksContainer from "../components/UnassignedBooksContainer";

export default function Home() {
  return (
    <>
      {/*Show a list of shelves with names*/}
      <ShelvesPreviewContainer />

      {/*Show list of currently unassigned books if any*/}

      <UnassignedBooksContainer />
    </>
  );
}
